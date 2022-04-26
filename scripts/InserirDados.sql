INSERT INTO minhascriptosprincipal.usuario (nome,email,senha) VALUES
	 ('Docker Test Um','docker@gmail.com','123456'),
	 ('Docker Test Dois','docker2@gmail.com','123456'),
	 ('Docker Test Tres','docker3@gmail.com','123456'),
	 ('Docker Test Quatro','docker4@gmail.com','123456');

INSERT INTO minhascriptosprincipal.criptomoeda ("tipoMoeda","dataDeCompra","quantidadeComprada","precoDeCompra","valorDaUnidadeNoDiaDeCompra",usuario_id) VALUES
('DOGE','2022-04-21',10.0,10.0,1.0,1),
('SHIBA','2022-04-24',110.0,110.0,0.01,1),
('BITCOIN','2022-04-21',0.00001,10.0,195000,1),
('ETH','2022-04-22',0.0001,10,14500,1),
('DOGE','2022-04-23',10.0,10.0,1.0,2),
('SHIBA','2022-04-22',110.0,110.0,0.01,2),
('BITCOIN','2022-04-22',0.00001,10.0,195000,2),
('ETH','2022-04-22',0.0001,10,14500,2),
('DOGE','2022-04-23',10.0,10.0,1.0,3),
('SHIBA','2022-04-22',110.0,110.0,0.01,3),
('BITCOIN','2022-04-22',0.00001,10.0,195000,3),
('ETH','2022-04-22',0.0001,10,14500,3);