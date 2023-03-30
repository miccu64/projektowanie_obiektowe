<?php

namespace App\Repository;

use App\Entity\Product;
use Doctrine\Bundle\DoctrineBundle\Repository\ServiceEntityRepository;
use Doctrine\Persistence\ManagerRegistry;

/**
 * @extends ServiceEntityRepository<Product>
 *
 * @method Product|null find($id, $lockMode = null, $lockVersion = null)
 * @method Product|null findOneBy(array $criteria, array $orderBy = null)
 * @method Product[]    findAll()
 * @method Product[]    findBy(array $criteria, array $orderBy = null, $limit = null, $offset = null)
 */
class ProductRepository extends ServiceEntityRepository
{
    public function __construct(ManagerRegistry $registry)
    {
        parent::__construct($registry, Product::class);
    }

    public function add(Product $entity): void
    {
        $this->getEntityManager()->persist($entity);
        $this->getEntityManager()->flush();
    }

    public function remove(int $id): void
    {
        $dbProduct = $this->getEntityManager()->getReference("App\\Entity\\Product", $id);
        $this->getEntityManager()->remove($dbProduct);
        $this->getEntityManager()->flush();
    }

    public function update(Product $product) {
        $dbProduct = $this->getEntityManager()->getReference("App\\Entity\\Product", $product->id);
        $dbProduct->name = $product->name;
        $dbProduct->price = $product->price;
        $this->getEntityManager()->flush();
    }
}
