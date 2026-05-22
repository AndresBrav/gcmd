package commands

// GitCommand representa un comando de Git individual
type GitCommand struct {
	Cmd       string // El comando en sí
	ShortDesc string // Descripción corta para la lista
	LongDesc  string // Descripción extendida explicativa
	Category  string // Categoría del comando
}

// Title implementa list.DefaultItem para mostrar el título en la TUI
func (g GitCommand) Title() string {
	return g.Cmd
}

// Description implementa list.DefaultItem para mostrar la descripción corta en la TUI
func (g GitCommand) Description() string {
	return "[" + g.Category + "] " + g.ShortDesc
}

// FilterValue implementa list.Item para permitir búsquedas difusas sobre el comando, descripción y categoría
func (g GitCommand) FilterValue() string {
	return g.Cmd + " " + g.ShortDesc + " " + g.Category
}

// GetCommands retorna la lista completa de comandos predefinidos
func GetCommands() []GitCommand {
	return []GitCommand{
		// Categoría: Configuración y Creación
		{
			Cmd:       "git init",
			ShortDesc: "Inicializar un repositorio Git local",
			LongDesc:  "Crea un nuevo repositorio local de Git en el directorio actual. Genera la carpeta oculta '.git'.",
			Category:  "Configuración y Creación",
		},
		{
			Cmd:       "git clone <url>",
			ShortDesc: "Clonar un repositorio existente",
			LongDesc:  "Descarga y copia un repositorio remoto existente desde la URL especificada a una carpeta local.",
			Category:  "Configuración y Creación",
		},
		{
			Cmd:       "git config --global user.name \"<nombre>\"",
			ShortDesc: "Configurar el nombre de usuario global",
			LongDesc:  "Establece el nombre que se asociará con todos tus commits en esta máquina de forma global.",
			Category:  "Configuración y Creación",
		},
		{
			Cmd:       "git config --global user.email \"<correo>\"",
			ShortDesc: "Configurar el correo electrónico global",
			LongDesc:  "Establece el correo electrónico que se asociará con todos tus commits en esta máquina de forma global.",
			Category:  "Configuración y Creación",
		},

		// Categoría: Guardado y Estado
		{
			Cmd:       "git status",
			ShortDesc: "Mostrar el estado actual del área de trabajo",
			LongDesc:  "Lista los archivos modificados, agregados o sin rastrear en comparación con el último commit.",
			Category:  "Guardado y Estado",
		},
		{
			Cmd:       "git add .",
			ShortDesc: "Preparar todos los cambios del directorio",
			LongDesc:  "Agrega todos los archivos nuevos, modificados y eliminados del directorio actual al área de preparación (staging).",
			Category:  "Guardado y Estado",
		},
		{
			Cmd:       "git add <archivo>",
			ShortDesc: "Preparar un archivo específico",
			LongDesc:  "Agrega los cambios de un archivo o directorio específico al área de preparación (staging) para el próximo commit.",
			Category:  "Guardado y Estado",
		},
		{
			Cmd:       "git commit -m \"<mensaje>\"",
			ShortDesc: "Confirmar los cambios preparados",
			LongDesc:  "Guarda los cambios del área de preparación en el historial del repositorio local con un mensaje explicativo.",
			Category:  "Guardado y Estado",
		},
		{
			Cmd:       "git commit --amend",
			ShortDesc: "Modificar el último commit",
			LongDesc:  "Permite editar el mensaje del último commit o agregar archivos que olvidaste incluir en él.",
			Category:  "Guardado y Estado",
		},
		{
			Cmd:       "git stash",
			ShortDesc: "Guardar temporalmente los cambios no confirmados",
			LongDesc:  "Guarda temporalmente el trabajo sin confirmar (modificados y preparados) en una pila interna para poder cambiar de rama.",
			Category:  "Guardado y Estado",
		},
		{
			Cmd:       "git stash pop",
			ShortDesc: "Recuperar los cambios guardados en stash",
			LongDesc:  "Extrae y vuelve a aplicar el conjunto de cambios guardado más recientemente en la pila del 'stash'.",
			Category:  "Guardado y Estado",
		},

		// Categoría: Ramas y Fusión
		{
			Cmd:       "git branch",
			ShortDesc: "Listar ramas locales",
			LongDesc:  "Muestra todas las ramas creadas localmente. La rama en la que te encuentras estará marcada con un asterisco (*).",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git branch -a",
			ShortDesc: "Listar todas las ramas (locales y remotas)",
			LongDesc:  "Muestra todas las ramas locales del repositorio y también las ramas de seguimiento remoto que han sido descargadas.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git branch <nombre>",
			ShortDesc: "Crear una nueva rama",
			LongDesc:  "Crea una rama apuntando al commit actual con el nombre especificado, pero no te cambia a ella automáticamente.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git checkout <rama>",
			ShortDesc: "Cambiar a una rama o commit",
			LongDesc:  "Actualiza los archivos de tu directorio de trabajo para que coincidan con la rama o commit especificado.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git checkout -b <nueva-rama>",
			ShortDesc: "Crear una rama y cambiar a ella inmediatamente",
			LongDesc:  "Crea una nueva rama local y cambia tu directorio de trabajo a ella en un solo paso rápido.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git switch <rama>",
			ShortDesc: "Cambiar a una rama (Alternativa moderna)",
			LongDesc:  "Comando moderno recomendado para cambiar de rama de forma segura en Git.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git switch -c <nueva-rama>",
			ShortDesc: "Crear y cambiar de rama (Alternativa moderna)",
			LongDesc:  "Comando moderno recomendado para crear una nueva rama y cambiarte a ella de inmediato.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git merge <rama>",
			ShortDesc: "Fusionar una rama en la rama actual",
			LongDesc:  "Combina el historial de commits y cambios de la rama especificada dentro de tu rama actual.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git rebase <rama>",
			ShortDesc: "Reorganizar tus commits sobre otra rama",
			LongDesc:  "Toma los commits locales de tu rama actual y los vuelve a aplicar uno a uno arriba de la rama especificada.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git branch -d <rama>",
			ShortDesc: "Eliminar una rama de forma segura",
			LongDesc:  "Borra la rama local especificada. Git impedirá que se elimine si contiene cambios no fusionados.",
			Category:  "Ramas y Fusión",
		},
		{
			Cmd:       "git branch -D <rama>",
			ShortDesc: "Forzar la eliminación de una rama",
			LongDesc:  "Elimina la rama local especificada sin importar si sus cambios han sido completamente fusionados o no.",
			Category:  "Ramas y Fusión",
		},

		// Categoría: Historial e Inspección
		{
			Cmd:       "git log --oneline",
			ShortDesc: "Ver historial simplificado en una línea",
			LongDesc:  "Muestra una lista condensada de commits con su ID abreviado y el mensaje de confirmación.",
			Category:  "Historial e Inspección",
		},
		{
			Cmd:       "git log --graph --oneline --all",
			ShortDesc: "Ver historial en formato de árbol gráfico",
			LongDesc:  "Dibuja una representación visual ASCII con las líneas temporales de todas las ramas locales y remotas.",
			Category:  "Historial e Inspección",
		},
		{
			Cmd:       "git diff",
			ShortDesc: "Diferencias del directorio de trabajo",
			LongDesc:  "Muestra los cambios detallados línea por línea que has hecho en tus archivos y que aún no has agregado al staging.",
			Category:  "Historial e Inspección",
		},
		{
			Cmd:       "git diff --staged",
			ShortDesc: "Diferencias de los cambios listos en staging",
			LongDesc:  "Muestra los cambios detallados entre el área de preparación (staging) y la última confirmación guardada (HEAD).",
			Category:  "Historial e Inspección",
		},
		{
			Cmd:       "git show <commit>",
			ShortDesc: "Mostrar los detalles de un commit específico",
			LongDesc:  "Muestra la descripción completa del commit y el parche (diff) con los cambios de código correspondientes.",
			Category:  "Historial e Inspección",
		},

		// Categoría: Sincronización
		{
			Cmd:       "git remote -v",
			ShortDesc: "Listar repositorios remotos configurados",
			LongDesc:  "Muestra las URLs de los servidores remotos conectados (generalmente 'origin') para subir o bajar código.",
			Category:  "Sincronización",
		},
		{
			Cmd:       "git fetch",
			ShortDesc: "Descargar novedades del servidor remoto",
			LongDesc:  "Trae el historial e información nueva del repositorio remoto a tu base de datos local sin modificar tu código.",
			Category:  "Sincronización",
		},
		{
			Cmd:       "git pull",
			ShortDesc: "Descargar y fusionar cambios remotos",
			LongDesc:  "Descarga los commits remotos de la rama actual y los fusiona de inmediato con tu rama local de trabajo.",
			Category:  "Sincronización",
		},
		{
			Cmd:       "git push origin <rama>",
			ShortDesc: "Subir commits locales al remoto",
			LongDesc:  "Envía tus cambios confirmados localmente a la rama correspondiente en el repositorio remoto 'origin'.",
			Category:  "Sincronización",
		},

		// Categoría: Correcciones y Reversiones
		{
			Cmd:       "git reset --soft HEAD~1",
			ShortDesc: "Deshacer el último commit conservando cambios",
			LongDesc:  "Deshace el último commit. Los archivos modificados se quedan en el área de preparación (staging) intactos.",
			Category:  "Correcciones y Reversiones",
		},
		{
			Cmd:       "git reset --hard HEAD~1",
			ShortDesc: "Deshacer el último commit borrando cambios",
			LongDesc:  "Deshace el último commit y destruye absolutamente todas las modificaciones de código hechas en ese commit. ¡Cuidado!",
			Category:  "Correcciones y Reversiones",
		},
		{
			Cmd:       "git revert <commit>",
			ShortDesc: "Deshacer los cambios de un commit mediante otro commit",
			LongDesc:  "Crea un nuevo commit seguro que hace exactamente lo opuesto al commit especificado, ideal para ramas públicas.",
			Category:  "Correcciones y Reversiones",
		},
		{
			Cmd:       "git cherry-pick <commit>",
			ShortDesc: "Aplicar un commit específico de otra rama",
			LongDesc:  "Toma la confirmación indicada de cualquier otra rama y la copia y aplica como un nuevo commit en tu rama actual.",
			Category:  "Correcciones y Reversiones",
		},
		{
			Cmd:       "git clean -fd",
			ShortDesc: "Eliminar archivos y carpetas sin rastrear",
			LongDesc:  "Limpia tu espacio de trabajo borrando permanentemente todos los archivos (-f) y carpetas (-d) que no estén en el control de Git.",
			Category:  "Correcciones y Reversiones",
		},
	}
}
