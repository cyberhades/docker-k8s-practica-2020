## Práctica 2020
 
El objetivo de esta práctica es ejecutar la aplicación que se encuentra en este mismo repositorio en dos entornos distintos: usando docker-compose y en un cluster de Kubernetes.

La aplicación adjunta es muy sencilla, es simplemente una aplicación web, que escucha por el puerto 8080 y va contando el número de visitas a la misma. Para guardar dicho número usa una base de datos tipo Redis versión 6.

Además de la aplicación, se adjuntan varios ficheros más, entre ellos un fichero Dockerfile que compila dicha aplicación y crea una imagen docker.

Se piden varias cosas:


### 1. Modificación del fichero Dockerfile

El fichero Dockerfile adjunto crea una imagen con muchas cosas que la aplicación no necesita en ejecución. Por lo tanto se pide que se actualice dicho fichero en un fichero multifase, cuya imagen final esté basada en: gcr.io/distroless/base

Además de eso, se pide que la aplicación se ejecute con un usuario llamado `app` cuyo id sea 1000.

### 2. Terminar el fichero docker-compose.yaml

Como también se puede observar, en el repositorio existe un fichero docker-compose.yaml incompleto con dos servicios definidos: web y redis.

Se pide que ambos servicios, web y redis, se ejecuten en una red llamada `backend`, y que el servicio `redis` se arranque antes que `web`.

El servicio `web` se conecta a la base de datos (redis) especificadad a través de la variable de entorno: REDIS_SERVER.

### 3. Crear los ficheros **manifesto** de Kubernetes para el servicio web.

En el repositorio también tenéis los ficheros para la creación del despliegue y el servicio para la base de datos Redis.

En esta parte de la práctica se pide la creación del despliegue con una copia (replica) del mismo, así como un servicio para exponer la aplicación y poder acceder desde afuera del clúster.

En este punto habréis podido comprobar que la base de datos no tiene habilitada ningún tipo de autenticación, por lo que cualquier pod del cluster podría conectarse a la base de datos y borrar o modificar el contador de visitas.
Por ello se pide que para evitar eso, se proteja la base de datos con políticas de red. De hecho se pide que se deniegue todo el tráfico del espacio de nombres y que luego se creen las políticas para abrir las conexiones necesarias.

## Notas

Para llevar a cabo esta práctica sobre todo la parte de políticas de red de Kubernetes, necesitas que tu cluster soporta las mismas. 
Si usas `minikube >= v1.12.1`, debes de arrancarlo de la siguiente manera:

    minikube start --cni=cilium --memory=4096

Si usas una versión anterior:

    minikube start --network-plugin=cni --memory=4096

    minikube ssh -- sudo mount bpffs -t bpf /sys/fs/bpf

    kubectl create -f https://raw.githubusercontent.com/cilium/cilium/v1.9/install/kubernetes/quick-install.yaml

Para más información con respecto cilium en minikube: https://docs.cilium.io/en/v1.9/gettingstarted/minikube/

Si lo prefieres, puedes usar otro cluster o CNI plugin.

#### Fortificando el contenedor en Kubernetes

Si quieres sacar puntuación extra, se pide que se fortifique el contenedor de la aplicación (no de Redis) de la siguiente manera:

- Forzando que el contenedor no se ejecute como root
- Montando el sistema de ficheros como sólo lectura
- No montando el token de la cuenta de servición asociada
- No permitiendo la escalada de privilegios del contenedor
- Sólo añadiendo las capacidades necesarias, si se necesitara alguna

## Puntuación de la práctica

- Creación del Dockerfile multifase - 2 puntos
- Creación correcta del fihero docker-compose - 2 puntos
- Creación del fichero de despliegue y servicio de la aplicación - 2 puntos
- Creación de las políticas de red - 2 puntos
- Fortificación del contenedor de la aplicación - 2 puntos