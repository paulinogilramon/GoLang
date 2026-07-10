Fase 1 - Fundamentos
Día 1
Variables
Constantes
Tipos básicos
if
for
switch

Proyecto: Calculadora.

Día 2
Funciones
Múltiples valores de retorno
Parámetros
Closures

Proyecto: Biblioteca de utilidades matemáticas.

Fase 2 - Datos
Día 3
Arrays
Slices
append
copy
range

Proyecto: Gestor de lista de compras.

Día 4
Maps
Operaciones sobre maps
delete
Comprobación de existencia

Proyecto: Agenda telefónica.

Fase 3 - Programación orientada a objetos (al estilo Go)
Día 5
Structs
Métodos
Receivers por valor
Receivers por puntero

Proyecto: Sistema de empleados.

Día 6
Embedding
Interfaces
Polimorfismo

Proyecto: Simulador de animales.

Fase 4 - Errores
Día 7
Tipo error
errors.New
fmt.Errorf
panic
recover
defer

Proyecto: Lector de archivos robusto.

Fase 5 - Concurrencia
Día 8
Goroutines
sync.WaitGroup

Proyecto: Descargar varios archivos en paralelo (simulado).

Día 9
Channels
Channels con buffer
select

Proyecto: Sistema productor-consumidor.

Fase 6 - Genéricos
Día 10
Type parameters
Restricciones (constraints)
Funciones genéricas
Tipos genéricos

Proyecto: Colección genérica (Stack[T]).

Fase 7 - Organización
Día 11
Paquetes
Módulos
Imports
Organización del código

Proyecto: Dividir proyectos anteriores en paquetes.

Fase 8 - Biblioteca estándar
Día 12

Aprender las librerías que más se usan:

fmt
os
io
bufio
strings
strconv
time
math
encoding/json
net/http

Proyecto: Cliente REST sencillo.

Fase 9 - Proyecto completo
Días 13 y 14

Crear un gestor de tareas.

Que tenga:

CRUD
JSON
Interfaces
Structs
Métodos
Slices
Maps
Errores
Paquetes
Goroutines
Channels
Genéricos (si tienen sentido)
defer
select
Cosas que aprendería después (no son parte del lenguaje)

Cuando ya domines el lenguaje, pasaría a su ecosistema:

Testing (testing)
Benchmarks
Profiling
Context (context.Context)
HTTP (net/http)
Concurrencia avanzada (sync.Mutex, sync.RWMutex, sync.Once, sync.Cond)
Reflexión (reflect)
Generación de código (go generate)
Compilación cruzada
Herramientas (go fmt, go vet, go test, go mod)
¿Y eso es todo?

Sí. Una de las cosas más llamativas de Go es que el lenguaje en sí es sorprendentemente pequeño. Si completas este plan, habrás aprendido prácticamente todas las características principales del lenguaje. A partir de ahí, la diferencia entre un desarrollador principiante y uno experto no suele estar en conocer más sintaxis, sino en escribir código más claro, organizar mejor los paquetes y sacar partido de la biblioteca estándar y del ecosistema.