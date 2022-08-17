# Prueba tecnica - Poker

## Tech Stack

**Database:** PostgreSQL

**Backend:** Golang


##

Running Tests

```bash
  go test
```

¿Cuál es el jugador más viejo de cada equipo?

```sql
SELECT queryT.edad, jugadores.nombre, equipos.nombre FROM (
  	SELECT fk_equipos, min(fecha_nacimiento) as edad FROM jugadores GROUP BY fk_equipos
	) AS queryT
 	INNER JOIN jugadores on queryT.edad = jugadores.fecha_nacimiento
 	INNER JOIN equipos on queryT.fk_equipos = equipos.id_equipos
 	ORDER BY queryT.edad ASC
```

¿Cuántos partidos jugó de visitante cada equipo? (nota: hay equipos no jugaron ningún partido)

```sql
SELECT equipos.nombre, count(fk_equipo_visitante) as cantidad_juegos_visitantes 
	FROM partidos
 	INNER JOIN  equipos ON partidos.fk_equipo_visitante = equipos.id_equipos
 	GROUP BY fk_equipo_visitante , equipos.id_equipos
```

¿Qué equipos jugaron el 01/01/2016 y el 12/02/2016?
```sql
SELECT  fk_equipo_local as equipos, equipos.nombre  FROM partidos
	INNER JOIN  
		equipos ON partidos.fk_equipo_local = equipos.id_equipos
	WHERE fecha_partido BETWEEN '2016-01-01' AND '2016-12-02'

UNION

SELECT  fk_equipo_visitante as equipos, equipos.nombre FROM partidos 
	INNER JOIN  
		equipos on partidos.fk_equipo_visitante = equipos.id_equipos
	WHERE fecha_partido BETWEEN '2016-01-01' AND '2016-12-02'

ORDER BY equipos;
```

Diga el total de goles que hizo el equipo “Chacarita” en su historia (como local o visitante)

```sql
SELECT SUM(total) as total_goles, nombre FROM (
	SELECT count(goles_local) as total, equipos.nombre as nombre
 	FROM partidos
 		INNER JOIN  equipos on partidos.fk_equipo_local = equipos.id_equipos
 	WHERE equipos.nombre = 'Chacarita' and goles_local <> 0
 	GROUP BY equipos.nombre
 
UNION ALL 

SELECT count(goles_visitante) as total, equipos.nombre as nombre
	FROM partidos
  		INNER JOIN  equipos on partidos.fk_equipo_visitante = equipos.id_equipos
 	WHERE equipos.nombre = 'Chacarita' and goles_visitante <> 0
  	GROUP BY equipos.nombre
) as temporal  GROUP BY nombre
```
    
## Preguntas

Supone que en un repositorio GIT hiciste un commit y te olvidaste un archivo.
Explicar como se soluciona si hiciste push, y como si aun no hiciste. De ser posible, buscar que quede solo
un commit con los cambios.

Si se realizo el commit pero aun no se hace el push a la rama, se realiza los siguientes pasos


```bash
--amend
```

 Nos permite combinar cambios por etapas con el commit anterior en lugar de crear un commit nuevo

 Ej.

```bash
git add archivo_omitido.txt
git commit --amend
```
Si se realizo el commit y se le hizo push hay que realizar el siguiente comando.


```bash
git reset [--mixed] HEAD~1
```
La cual matendra los cambios.

Tenes un sitio en tu computadora de desarrollo, y cuando entras desde el navegador, en la consola te
aparece esto:


```bash
https://site.local/favicon.ico Failed to load resource: net::ERR_INSECURE_RESPONSE
```

El archivo existe, el sitio debe cargar por https. Como lo solucionas?

Las principales soluciones para este caso serian: 

- Agregar como un sitio de confianza, solamente si se esta trabajando con el dominion en local
- Actualizar la configuración de fecha / hora del dispositivo; de alguna manera la fecha no se estableció correctamente.


Tienes un archivo comiteado en un repositorio GIT que deseas que quede ignorado. Que pasos debes
realizar?

Se tiene que crear un archivo llamado **.gitignore** luego colocar el nombre del archivo que deseamos ignorar a partir de ahora.

Explica las ventajas de cargar en tu sitio las librerias de terceros por GTM.

- Debido a que todo está en un solo lugar, es más fácil solucionar y corregir errores de etiquetas, incluso antes de que se publiquen
- Es absolutamente gratuito y es perfecto para pequeñas y medianas empresas.
- No se requieren conocimientos de programación. Casi cualquier persona puede realizar actualizaciones fácilmente, agregar nuevas etiquetas, probar cada cambio e implementar etiquetas sin necesidad de realizar una codificación compleja del sitio web.


Contanos en pocas lineas cual crees que es la mayor innovacion en el mundo del desarrollo en los ultimos 5
años, y por que.

La mayor innovacion de estos ultimos a;os ha sido GitHub Copilot resulta ser emocionante para la comunidad de desarrolladores porque promete que puede reducir la barrera de la codificación, lo que nos ayudaría muchisimo procesar y entregar códigos en un tiempo menor.