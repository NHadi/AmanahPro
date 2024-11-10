CREATE TABLE Users (
    UserId INT PRIMARY KEY IDENTITY,
    Username NVARCHAR(100) NOT NULL UNIQUE,
    PasswordHash NVARCHAR(255) NOT NULL,
    Email NVARCHAR(255),
    CreatedAt DATETIME DEFAULT GETDATE(),
    IsActive BIT DEFAULT 1
);

CREATE TABLE Roles (
    RoleId INT PRIMARY KEY IDENTITY,
    RoleName NVARCHAR(100) NOT NULL UNIQUE,
    Description NVARCHAR(255)
);

CREATE TABLE Permissions (
    PermissionId INT PRIMARY KEY IDENTITY,
    PermissionName NVARCHAR(100) NOT NULL UNIQUE,
    Description NVARCHAR(255)
);

CREATE TABLE Menus (
    MenuId INT PRIMARY KEY IDENTITY,
    MenuName NVARCHAR(100) NOT NULL,
    Path NVARCHAR(255) NOT NULL,
    ParentId INT NULL,  -- For hierarchical menus
    FOREIGN KEY (ParentId) REFERENCES Menus(MenuId)
);

CREATE TABLE RolePermissions (
    RoleId INT,
    PermissionId INT,
    PRIMARY KEY (RoleId, PermissionId),
    FOREIGN KEY (RoleId) REFERENCES Roles(RoleId),
    FOREIGN KEY (PermissionId) REFERENCES Permissions(PermissionId)
);

CREATE TABLE UserRoles (
    UserId INT,
    RoleId INT,
    PRIMARY KEY (UserId, RoleId),
    FOREIGN KEY (UserId) REFERENCES Users(UserId),
    FOREIGN KEY (RoleId) REFERENCES Roles(RoleId)
);

CREATE TABLE MenuPermissions (
    MenuId INT,
    PermissionId INT,
    PRIMARY KEY (MenuId, PermissionId),
    FOREIGN KEY (MenuId) REFERENCES Menus(MenuId),
    FOREIGN KEY (PermissionId) REFERENCES Permissions(PermissionId)
);
