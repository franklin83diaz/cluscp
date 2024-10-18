# cluscp

## Descripción

**cluscp** (Copiar en Cluster) es una herramienta desarrollada para copiar grandes cantidades de datos de manera eficiente hacia múltiples nodos de un clúster HPC. Usa una técnica de copiado en cascada que tiene un rendimiento similar a usar multicast, pero sin la necesidad de que esté habilitado multicast en la red de nodos.

## Características

- Un programa simple y portable (disponible para instalación automática).
- Copia archivos de una ubicación a otra de manera rápida y eficiente.
- Permite copiar hacia nodos que estén dentro de una red privada.
- No necesita ser instalado manualmente en los nodos.
- No requiere conexión a servidores externos.
- No requiere multicast o broadcast.
- Opción de transmisión de datos comprimidos.
- Transmisión de datos cifrada (cuando se usa conexión SSL).

## Instalación

Para instalar **cluscp**, sigue estos pasos:

1. **Descargar cluscp**: Clona el repositorio desde GitHub o descarga el archivo binario precompilado.

2. **Instalación en el cliente o desintalar**: 

   `cluscp install`  or  `cluscp unistall`

3. **Instalación en los nodos**: Para instalar **cluscp** en los nodos del clúster, utiliza el archivo de configuración `nodos.conf` que contiene la lista de nodos de la red privada.

   Nota: La instalación en los nodos no es obligatoria, pero puede ahorrar tiempo durante el copiado.

   Nota: el primero host en el nodo de conexion y se usa para distrubuir a la red privada en caso de.

## Uso

La sintaxis básica para utilizar **cluscp** es la siguiente:

### Ejemplos

- Para copiar un archivo simple:
