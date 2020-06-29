use ecommerce;

CREATE TABLE `produtos`(
    `id` bigint(10) NOT NULL AUTO_INCREMENT,
    `nome` varchar(60) NOT NULL,
    `descricao` varchar(60) NOT NULL,
    `preco` decimal(10,2) NOT NULL,
    `quantidade` bigint(10) NOT NULL,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;

insert into produtos (nome, descricao, preco, quantidade) values
('Camiseta', 'Azul Marinho', 39.0, 5),
('Tenis', 'Vermelho Sangue', 105.1, 5),
('Meia', 'Verde Mato', 15.9, 5);