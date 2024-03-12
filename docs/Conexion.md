1. Se inicializa el server corriendo main.go
2. El server deja abierto el puerto :8080 sirviendo la ruta /ws. Se queda esperando a que alguien solicite conexion en ese puerto.
3. Una vez que alguien solicita conexion. Corre "serveWebSocket". Incializando el cliente e insertandolo en el pool de clientes. Luego se queda escuchando mensajes de la connexion "conn".