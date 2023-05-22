import Fluent

struct CreateProducts: AsyncMigration {
    func prepare(on database: Database) async throws {
        try await database.schema("products")
            .id()
            .field("name", .string)
            .field("price", .double)
            .create()
    }

    func revert(on database: Database) async throws {
        try await database.schema("products").delete()
    }
}