<?php

namespace App\Controller;

use App\Entity\Car;
use App\Repository\CarRepository;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Annotation\Route;

class CarController extends AbstractController
{
    public function __construct(
        private CarRepository $carRepository,
    )
    {
    }

    #[Route('/cars', methods: ['GET', 'HEAD'])]
    public function getAllCars(): JsonResponse
    {
        return new JsonResponse($this->carRepository->findAll());
    }

    #[Route('/cars/{id}', methods: ['GET', 'HEAD'])]
    public function getCar(int $id): JsonResponse
    {
        return new JsonResponse($this->carRepository->find($id));
    }

    #[Route('/cars', methods: ['POST'])]
    public function addCar(Request $request): Response
    {
        try {
            $car = $this->jsonDecode($request);

            $this->carRepository->add($car);
            return new Response(
                "{\"id\":" . $car->id . "}",
                Response::HTTP_OK
            );
        } catch (\Exception $exception) {
            return new Response(
                $exception->getMessage(),
                Response::HTTP_NOT_ACCEPTABLE
            );
        }
    }

    #[Route('/cars', methods: ['PATCH'])]
    public function patchCar(Request $request): Response
    {
        try {
            $car = $this->jsonDecode($request);
            $this->carRepository->update($car);
            return new Response(
                "{\"id\":" . $car->id . "}",
                Response::HTTP_OK
            );
        } catch (\Exception $exception) {
            return new Response(
                $exception->getMessage(),
                Response::HTTP_NOT_ACCEPTABLE
            );
        }
    }

    #[Route('/cars/{id}', methods: ['DELETE'])]
    public function deleteCar(int $id): Response
    {
        try {
            $this->carRepository->remove($id);
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

    private function jsonDecode(Request $request): Car
    {
        $data = json_decode($request->getContent(), true);
        $car = new Car();
        foreach ($data as $key => $value) $car->{$key} = $value;
        return $car;
    }
}