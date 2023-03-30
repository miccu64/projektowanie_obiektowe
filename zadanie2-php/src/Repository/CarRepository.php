<?php

namespace App\Repository;

use App\Entity\Car;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends ServiceEntityRepository<Car>
 *
 * @method Car|null find($id, $lockMode = null, $lockVersion = null)
 * @method Car|null findOneBy(array $criteria, array $orderBy = null)
 * @method Car[]    findAll()
 * @method Car[]    findBy(array $criteria, array $orderBy = null, $limit = null, $offset = null)
 */
class CarRepository extends ServiceEntityRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Car::class);
    }

    public function add(Car $entity): void
    {
        $this->getEntityManager()->persist($entity);
        $this->getEntityManager()->flush();
    }

    public function remove(int $id): void
    {
        $dbCar = $this->getEntityManager()->getReference("App\\Entity\\Car", $id);
        $this->getEntityManager()->remove($dbCar);
        $this->getEntityManager()->flush();
    }

    public function update(Car $car) {
        $dbCar = $this->getEntityManager()->getReference("App\\Entity\\Car", $car->id);
        $dbCar->color = $car->color;
        $dbCar->type = $car->type;
        $dbCar->isSuv = $car->isSuv;
        $this->getEntityManager()->flush();
    }
}
