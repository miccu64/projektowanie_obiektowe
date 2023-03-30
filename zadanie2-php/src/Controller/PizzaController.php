<?php

namespace App\Controller;

use App\Entity\Pizza;
use App\Repository\PizzaRepository;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

class PizzaController extends AbstractController
{
    public function __construct(
        private PizzaRepository $pizzaRepository,
    )
    {
    }

    #[Route('/pizzas', methods: ['GET', 'HEAD'])]
    public function getAllPizzas(): JsonResponse
    {
        return new JsonResponse($this->pizzaRepository->findAll());
    }

    #[Route('/pizzas/{id}', methods: ['GET', 'HEAD'])]
    public function getPizza(int $id): JsonResponse
    {
        return new JsonResponse($this->pizzaRepository->find($id));
    }

    #[Route('/pizzas', methods: ['POST'])]
    public function addPizza(Request $request): Response
    {
        try {
            $pizza = $this->jsonDecode($request);

            $this->pizzaRepository->add($pizza);
            return new Response(
                "{\"id\":" . $pizza->id . "}",
                Response::HTTP_OK
            );
        } catch (\Exception $exception) {
            return new Response(
                $exception->getMessage(),
                Response::HTTP_NOT_ACCEPTABLE
            );
        }
    }

    #[Route('/pizzas', methods: ['PATCH'])]
    public function patchPizza(Request $request): Response
    {
        try {
            $pizza = $this->jsonDecode($request);
            $this->pizzaRepository->update($pizza);
            return new Response(
                "{\"id\":" . $pizza->id . "}",
                Response::HTTP_OK
            );
        } catch (\Exception $exception) {
            return new Response(
                $exception->getMessage(),
                Response::HTTP_NOT_ACCEPTABLE
            );
        }
    }

    #[Route('/pizzas/{id}', methods: ['DELETE'])]
    public function deletePizza(int $id): Response
    {
        try {
            $this->pizzaRepository->remove($id);
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

    private function jsonDecode(Request $request): Pizza
    {
        $data = json_decode($request->getContent(), true);
        $pizza = new Pizza();
        foreach ($data as $key => $value) $pizza->{$key} = $value;
        return $pizza;
    }
}