<?php

namespace App\Controller;

use App\Entity\Product;
use App\Repository\ProductRepository;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

class ProductController extends AbstractController
{
    public function __construct(
        private ProductRepository $productRepository,
    )
    {
    }

    #[Route('/products', methods: ['GET', 'HEAD'])]
    public function getAllProducts(): JsonResponse
    {
        return new JsonResponse($this->productRepository->findAll());
    }

    #[Route('/products/{id}', methods: ['GET', 'HEAD'])]
    public function getProduct(int $id): JsonResponse
    {
        return new JsonResponse($this->productRepository->find($id));
    }

    #[Route('/products', methods: ['POST'])]
    public function addProduct(Request $request): Response
    {
        try {
            $product = $this->jsonDecode($request);

            $this->productRepository->add($product);
            return new Response(
                "{\"id\":" . $product->id . "}",
                Response::HTTP_OK
            );
        } catch (\Exception $exception) {
            return new Response(
                $exception->getMessage(),
                Response::HTTP_NOT_ACCEPTABLE
            );
        }
    }

    #[Route('/products', methods: ['PATCH'])]
    public function patchProduct(Request $request): Response
    {
        try {
            $product = $this->jsonDecode($request);
            $this->productRepository->update($product);
            return new Response(
                "{\"id\":" . $product->id . "}",
                Response::HTTP_OK
            );
        } catch (\Exception $exception) {
            return new Response(
                $exception->getMessage(),
                Response::HTTP_NOT_ACCEPTABLE
            );
        }
    }

    #[Route('/products/{id}', methods: ['DELETE'])]
    public function deleteProduct(int $id): Response
    {
        try {
            $this->productRepository->remove($id);
            return new Response(
                "{result: true}",
                Response::HTTP_OK
            );
        } catch (\Exception $exception) {
            return new Response(
                $exception->getMessage(),
                Response::HTTP_BAD_REQUEST
            );
        }
    }

    private function jsonDecode(Request $request): Product
    {
        $data = json_decode($request->getContent(), true);
        $product = new Product();
        foreach ($data as $key => $value) $product->{$key} = $value;
        return $product;
    }
}