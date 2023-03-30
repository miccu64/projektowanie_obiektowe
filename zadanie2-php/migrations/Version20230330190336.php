<?php

declare(strict_types=1);

namespace DoctrineMigrations;

use Doctrine\DBAL\Schema\Schema;
use Doctrine\Migrations\AbstractMigration;

/**
 * Auto-generated Migration: Please modify to your needs!
 */
final class Version20230330190336 extends AbstractMigration
{
    public function getDescription(): string
    {
        return '';
    }

    public function up(Schema $schema): void
    {
        // this up() migration is auto-generated, please modify it to your needs
        $this->addSql('CREATE TABLE car (id INTEGER NOT NULL, type VARCHAR(255) NOT NULL, color VARCHAR(255) NOT NULL, is_suv BOOLEAN NOT NULL, PRIMARY KEY(id))');
        $this->addSql('CREATE TEMPORARY TABLE __temp__pizza AS SELECT id, type, cost FROM pizza');
        $this->addSql('DROP TABLE pizza');
        $this->addSql('CREATE TABLE pizza (id INTEGER NOT NULL, type VARCHAR(255) NOT NULL, cost INTEGER NOT NULL, PRIMARY KEY(id))');
        $this->addSql('INSERT INTO pizza (id, type, cost) SELECT id, type, cost FROM __temp__pizza');
        $this->addSql('DROP TABLE __temp__pizza');
        $this->addSql('CREATE TEMPORARY TABLE __temp__product AS SELECT id, name, price FROM product');
        $this->addSql('DROP TABLE product');
        $this->addSql('CREATE TABLE product (id INTEGER NOT NULL, name VARCHAR(255) NOT NULL, price INTEGER NOT NULL, PRIMARY KEY(id))');
        $this->addSql('INSERT INTO product (id, name, price) SELECT id, name, price FROM __temp__product');
        $this->addSql('DROP TABLE __temp__product');
    }

    public function down(Schema $schema): void
    {
        // this down() migration is auto-generated, please modify it to your needs
        $this->addSql('DROP TABLE car');
        $this->addSql('CREATE TEMPORARY TABLE __temp__pizza AS SELECT id, type, cost FROM pizza');
        $this->addSql('DROP TABLE pizza');
        $this->addSql('CREATE TABLE pizza (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, type VARCHAR(255) NOT NULL, cost INTEGER NOT NULL)');
        $this->addSql('INSERT INTO pizza (id, type, cost) SELECT id, type, cost FROM __temp__pizza');
        $this->addSql('DROP TABLE __temp__pizza');
        $this->addSql('CREATE TEMPORARY TABLE __temp__product AS SELECT id, name, price FROM product');
        $this->addSql('DROP TABLE product');
        $this->addSql('CREATE TABLE product (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, name VARCHAR(255) NOT NULL, price INTEGER NOT NULL)');
        $this->addSql('INSERT INTO product (id, name, price) SELECT id, name, price FROM __temp__product');
        $this->addSql('DROP TABLE __temp__product');
    }
}
