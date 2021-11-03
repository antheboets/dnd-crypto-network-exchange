CREATE LOGIN dndUser WITH PASSWORD = 'cB345678';
GO

CREATE DATABASE dndDb;
GO

Use dndDb;
GO

IF NOT EXISTS (SELECT * FROM sys.database_principals WHERE name = N'dndUser')
BEGIN
    CREATE USER dndUser FOR LOGIN dndUser
    EXEC sp_addrolemember N'db_owner', N'dndUser'
END;
GO
