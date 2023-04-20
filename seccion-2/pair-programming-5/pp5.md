**Ejercicio 3 Pair programming:** 

<aside>
💡 .El objetivo es simular un sistema que procesa múltiples mensajes en paralelo, donde cada mensaje tiene un "tipo" (por ejemplo, "email", "SMS", "notificación push") y un "contenido" (un string con el contenido del mensaje)

</aside>

**Sistema de cola de mensajes utilizando gorutinas y canales**
Crea un programa en Go que simule un sistema de cola de mensajes. El programa debe ser capaz de procesar múltiples mensajes en paralelo utilizando gorutinas y canales. Cada mensaje debe tener un "tipo" (por ejemplo, "email", "SMS", "notificación push") y un "contenido" (un string con el contenido del mensaje). Además, cada tipo de mensaje debe ser procesado por una gorutina específica para ese tipo, y cada gorutina debe procesar los mensajes en orden de llegada.

El programa debe permitir a los usuarios agregar mensajes a la cola y recibir confirmación de que el mensaje ha sido procesado por la gorutina correspondiente. También debe ser posible detener y reiniciar el procesamiento de mensajes para cada tipo de mensaje individualmente.