CREATE TABLE IF NOT EXISTS Prioridades (
	Id SMALLINT PRIMARY KEY,
	Descripcion VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Roles (
	Id SMALLINT PRIMARY KEY,
	Descripcion VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Usuarios (
	Id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	Nombre VARCHAR(100) NOT NULL,
	RolId SMALLINT NOT NULL REFERENCES Roles(Id),
	EstaDesactivado BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS Tareas (
	Id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	Titulo VARCHAR(200) NOT NULL,
	Descripcion TEXT,
	PrioridadId SMALLINT NOT NULL REFERENCES Prioridades(Id),
	FechaEntrega TIMESTAMP,
	EstaTerminado BOOLEAN NOT NULL DEFAULT FALSE,
	EstaEliminado BOOLEAN NOT NULL DEFAULT FALSE,
	CreadoEl TIMESTAMP NOT NULL,
	CreadoPor INTEGER NOT NULL REFERENCES Usuarios(Id),
	ModificadoEl TIMESTAMP,
	ModificadoPor INTEGER REFERENCES Usuarios(Id)
);

CREATE TABLE IF NOT EXISTS SubTareas (
	Id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	TareaId BIGINT NOT NULL REFERENCES Tareas(Id) ON DELETE CASCADE,
	Titulo VARCHAR(200) NOT NULL,
	EstaTerminado BOOLEAN NOT NULL DEFAULT FALSE,
	EstaEliminado BOOLEAN NOT NULL DEFAULT FALSE,
	CreadoEl TIMESTAMP NOT NULL,
	CreadoPor INTEGER NOT NULL REFERENCES Usuarios(Id),
	ModificadoEl TIMESTAMP,
	ModificadoPor INTEGER REFERENCES Usuarios(Id)
);

CREATE TABLE IF NOT EXISTS TareaXUsuario (
	TareaId BIGINT NOT NULL REFERENCES Tareas(Id) ON DELETE CASCADE,
	UsuarioId INTEGER NOT NULL REFERENCES Usuarios(Id),
	PRIMARY KEY (TareaId, UsuarioId)
);

INSERT INTO Prioridades (Id, Descripcion)
VALUES
    (1, 'Urgente'),
    (2, 'Alto'),
    (3, 'Medio'),
    (4, 'Bajo')
ON CONFLICT (Id) DO NOTHING;

INSERT INTO Roles (Id, Descripcion)
VALUES
    (1, 'Lider'),
    (2, 'Contribuidor')
ON CONFLICT (Id) DO NOTHING;

INSERT INTO Usuarios (Nombre, RolId)
SELECT 'Daniel Rivera', 1
WHERE NOT EXISTS (
    SELECT 1
    FROM Usuarios
    WHERE Nombre = 'Daniel Rivera'
);

INSERT INTO Usuarios (Nombre, RolId)
SELECT 'Edi Torres', 2
WHERE NOT EXISTS (
    SELECT 1
    FROM Usuarios
    WHERE Nombre = 'Edi Torres'
);

-- INSERT INTO Usuarios (
-- 	Nombre, 
-- 	RolId
-- ) VALUES (
-- 	'Edi Torres',
-- 	2
-- );

-- INSERT INTO Tareas (
-- 	Titulo,
-- 	Descripcion,
-- 	PrioridadId,
-- 	FechaEntrega,
-- 	CreadoEl,
-- 	CreadoPor
-- )
-- VALUES (
-- 	'Test1',
-- 	'Estudiar GO, Postgre',
-- 	2,
-- 	'2026-09-30 18:00:00',
-- 	CURRENT_TIMESTAMP,
-- 	1
-- )
-- RETURNING Id;

-- INSERT INTO TareaXUsuario (
-- 	TareaId, 
-- 	UsuarioId
-- )
-- VALUES (
-- 	1, 
-- 	2
-- );