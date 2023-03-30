<?php

namespace App\Repository;

use App\Entity\Pizza;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends ServiceEntityRepository<Pizza>
 *
 * @method Pizza|null find($id, $lockMode = null, $lockVersion = null)
 * @method Pizza|null findOneBy(array $criteria, array $orderBy = null)
 * @method Pizza[]    findAll()
 * @method Pizza[]    findBy(array $criteria, array $orderBy = null, $limit = null, $offset = null)
 */
class PizzaRepository extends ServiceEntityRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Pizza::class);
    }

    public function add(Pizza $entity): void
    {
        $this->getEntityManager()->persist($entity);
        $this->getEntityManager()->flush();
    }

    public function remove(int $id): void
    {
        $dbPizza = $this->getEntityManager()->getReference("App\\Entity\\Pizza", $id);
        $this->getEntityManager()->remove($dbPizza);
        $this->getEntityManager()->flush();
    }

    public function update(Pizza $Pizza) {
        $dbPizza = $this->getEntityManager()->getReference("App\\Entity\\Pizza", $Pizza->id);
        $dbPizza->type = $Pizza->type;
        $dbPizza->cost = $Pizza->cost;
        $this->getEntityManager()->flush();
    }
}
