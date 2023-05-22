import Vapor

func routes(_ app: Application) throws {
    app.get("products") { req async throws in
        try await Product.query(on: req.db).all()
    }

    app.get("products", ":id") { req -> EventLoopFuture<Product> in
        Product.find(req.parameters.get("id"), on: req.db).unwrap(or: Abort(.notFound))
    }

    app.post("products") { req -> EventLoopFuture<Product> in
        let product = try req.content.decode(Product.self)
        return product.create(on: req.db)
                .map {
                    product
                }
    }

    app.delete("products", ":id") { req -> EventLoopFuture<HTTPStatus> in
        Product.find(req.parameters.get("id"), on: req.db)
                .unwrap(or: Abort(.notFound))
                .flatMap { product in
                    product.delete(on: req.db)
                            .transform(to: .noContent)
                }
    }

    app.put("products", ":id") { req -> EventLoopFuture<Product> in
        let updatedProduct = try req.content.decode(Product.self)
        return Product.find(
                        req.parameters.get("id"),
                        on: req.db)
                .unwrap(or: Abort(.notFound)).flatMap { product in
                    product.name = updatedProduct.name
                    product.price = updatedProduct.price
                    return product.save(on: req.db).map {
                        product
                    }
                }
    }
}
