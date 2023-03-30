<?php

namespace App\Entity;

use App\Repository\ProductRepository;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity(repositoryClass: ProductRepository::class)]
class Car
{
    #[ORM\Id]
    #[ORM\Column(type: 'integer')]
    public $id;

    #[ORM\Column(type: 'string', length: 255)]
    public $type;

    #[ORM\Column(type: 'string', length: 255)]
    public $color;

    #[ORM\Column(type: 'boolean')]
    public $isSuv;
}
