-- Users Table
CREATE TABLE Users (
    user_id UNIQUEIDENTIFIER DEFAULT NEWID() PRIMARY KEY,
    username NVARCHAR(50) UNIQUE NOT NULL,
    email NVARCHAR(100) UNIQUE NOT NULL,
    password NVARCHAR(255) NOT NULL,
    status BIT DEFAULT 1,
    created_at DATETIME DEFAULT GETDATE(),
    updated_at DATETIME DEFAULT GETDATE()
);

-- Roles Table
CREATE TABLE Roles (
    role_id UNIQUEIDENTIFIER DEFAULT NEWID() PRIMARY KEY,
    role_name NVARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    created_at DATETIME DEFAULT GETDATE()
);

-- UserRoles Table
CREATE TABLE UserRoles (
    user_id UNIQUEIDENTIFIER FOREIGN KEY REFERENCES Users(user_id) ON DELETE CASCADE,
    role_id UNIQUEIDENTIFIER FOREIGN KEY REFERENCES Roles(role_id) ON DELETE CASCADE,
    assigned_at DATETIME DEFAULT GETDATE(),
    PRIMARY KEY (user_id, role_id)
);

-- Menus Table
CREATE TABLE Menus (
    menu_id UNIQUEIDENTIFIER DEFAULT NEWID() PRIMARY KEY,
    parent_id UNIQUEIDENTIFIER FOREIGN KEY REFERENCES Menus(menu_id),
    menu_name NVARCHAR(50) NOT NULL,
    path NVARCHAR(100) NOT NULL,
    icon NVARCHAR(50),
    [order] INT,
    created_at DATETIME DEFAULT GETDATE()
);

-- RoleMenus Table
CREATE TABLE RoleMenus (
    role_id UNIQUEIDENTIFIER FOREIGN KEY REFERENCES Roles(role_id) ON DELETE CASCADE,
    menu_id UNIQUEIDENTIFIER FOREIGN KEY REFERENCES Menus(menu_id) ON DELETE CASCADE,
    permission NVARCHAR(10) CHECK (permission IN ('view', 'edit', 'delete')),
    assigned_at DATETIME DEFAULT GETDATE(),
    PRIMARY KEY (role_id, menu_id, permission)
);

go

-- Insert Seed Data for Roles
INSERT INTO Roles (role_id, role_name, description)
VALUES
    (NEWID(), 'Admin', 'Administrator with full access'),
    (NEWID(), 'Editor', 'User with editing rights'),
    (NEWID(), 'Viewer', 'User with view-only access');

-- Insert Seed Data for Users
INSERT INTO Users (user_id, username, email, password, status)
VALUES
    (NEWID(), 'admin_user', 'admin@example.com', 'hashed_password_1', 1),
    (NEWID(), 'editor_user', 'editor@example.com', 'hashed_password_2', 1),
    (NEWID(), 'viewer_user', 'viewer@example.com', 'hashed_password_3', 1);

-- Insert Seed Data for UserRoles
-- Linking Users to Roles
INSERT INTO UserRoles (user_id, role_id)
SELECT u.user_id, r.role_id FROM Users u, Roles r
WHERE u.username = 'admin_user' AND r.role_name = 'Admin';

INSERT INTO UserRoles (user_id, role_id)
SELECT u.user_id, r.role_id FROM Users u, Roles r
WHERE u.username = 'editor_user' AND r.role_name = 'Editor';

INSERT INTO UserRoles (user_id, role_id)
SELECT u.user_id, r.role_id FROM Users u, Roles r
WHERE u.username = 'viewer_user' AND r.role_name = 'Viewer';

-- Insert Seed Data for Menus
INSERT INTO Menus (menu_id, parent_id, menu_name, path, icon, [order])
VALUES
    (NEWID(), NULL, 'Dashboard', '/dashboard', 'dashboard_icon', 1),
    (NEWID(), NULL, 'Settings', '/settings', 'settings_icon', 2),
    (NEWID(), NULL, 'Reports', '/reports', 'reports_icon', 3),
    (NEWID(), NULL, 'Users', '/users', 'users_icon', 4),
    (NEWID(), NULL, 'Roles', '/roles', 'roles_icon', 5);

-- Insert Seed Data for RoleMenus
-- Admin has access to all menus with all permissions
INSERT INTO RoleMenus (role_id, menu_id, permission)
SELECT r.role_id, m.menu_id, 'view' FROM Roles r, Menus m WHERE r.role_name = 'Admin';
INSERT INTO RoleMenus (role_id, menu_id, permission)
SELECT r.role_id, m.menu_id, 'edit' FROM Roles r, Menus m WHERE r.role_name = 'Admin';
INSERT INTO RoleMenus (role_id, menu_id, permission)
SELECT r.role_id, m.menu_id, 'delete' FROM Roles r, Menus m WHERE r.role_name = 'Admin';

-- Editor has view and edit access to Dashboard and Reports
INSERT INTO RoleMenus (role_id, menu_id, permission)
SELECT r.role_id, m.menu_id, 'view' FROM Roles r, Menus m WHERE r.role_name = 'Editor' AND m.menu_name IN ('Dashboard', 'Reports');
INSERT INTO RoleMenus (role_id, menu_id, permission)
SELECT r.role_id, m.menu_id, 'edit' FROM Roles r, Menus m WHERE r.role_name = 'Editor' AND m.menu_name IN ('Dashboard', 'Reports');

-- Viewer has view-only access to Dashboard
INSERT INTO RoleMenus (role_id, menu_id, permission)
SELECT r.role_id, m.menu_id, 'view' FROM Roles r, Menus m WHERE r.role_name = 'Viewer' AND m.menu_name = 'Dashboard';
