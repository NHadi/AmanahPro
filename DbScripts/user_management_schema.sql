USE [AmanahDb]
GO
/****** Object:  Table [dbo].[Menus]    Script Date: 13/11/2024 22:40:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Menus](
	[menu_id] [int] IDENTITY(1,1) NOT NULL,
	[parent_id] [int] NULL,
	[name] [nvarchar](50) NOT NULL,
	[path] [nvarchar](100) NOT NULL,
	[icon] [nvarchar](50) NULL,
	[order] [int] NULL,
	[created_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[menu_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[RoleMenus]    Script Date: 13/11/2024 22:40:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[RoleMenus](
	[role_id] [int] NOT NULL,
	[menu_id] [int] NOT NULL,
	[permission] [varchar](10) NOT NULL,
	[assigned_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[role_id] ASC,
	[menu_id] ASC,
	[permission] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Roles]    Script Date: 13/11/2024 22:40:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Roles](
	[role_id] [int] IDENTITY(1,1) NOT NULL,
	[name] [varchar](100) NOT NULL,
	[description] [varchar](max) NULL,
	[created_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[role_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[UserRoles]    Script Date: 13/11/2024 22:40:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[UserRoles](
	[user_id] [int] NOT NULL,
	[role_id] [int] NOT NULL,
	[assigned_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[user_id] ASC,
	[role_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Users]    Script Date: 13/11/2024 22:40:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Users](
	[user_id] [int] IDENTITY(1,1) NOT NULL,
	[username] [varchar](100) NOT NULL,
	[email] [varchar](100) NOT NULL,
	[password] [varchar](100) NOT NULL,
	[status] [bit] NULL,
	[created_at] [datetime] NULL,
	[updated_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[user_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
SET IDENTITY_INSERT [dbo].[Menus] ON 
GO
INSERT [dbo].[Menus] ([menu_id], [parent_id], [name], [path], [icon], [order], [created_at]) VALUES (1, NULL, N'Dashboard', N'/dashboard', N'dashboard_icon', 1, CAST(N'2024-11-12T16:39:21.110' AS DateTime))
GO
INSERT [dbo].[Menus] ([menu_id], [parent_id], [name], [path], [icon], [order], [created_at]) VALUES (2, NULL, N'Settings', N'/settings', N'settings_icon', 2, CAST(N'2024-11-12T16:39:21.110' AS DateTime))
GO
INSERT [dbo].[Menus] ([menu_id], [parent_id], [name], [path], [icon], [order], [created_at]) VALUES (3, NULL, N'Reports', N'/reports', N'reports_icon', 3, CAST(N'2024-11-12T16:39:21.110' AS DateTime))
GO
INSERT [dbo].[Menus] ([menu_id], [parent_id], [name], [path], [icon], [order], [created_at]) VALUES (4, NULL, N'Users', N'/users', N'users_icon', 4, CAST(N'2024-11-12T16:39:21.110' AS DateTime))
GO
INSERT [dbo].[Menus] ([menu_id], [parent_id], [name], [path], [icon], [order], [created_at]) VALUES (5, NULL, N'Roles', N'/roles', N'roles_icon', 5, CAST(N'2024-11-12T16:39:21.110' AS DateTime))
GO
SET IDENTITY_INSERT [dbo].[Menus] OFF
GO
INSERT [dbo].[RoleMenus] ([role_id], [menu_id], [permission], [assigned_at]) VALUES (1, 1, N'edit', CAST(N'2024-11-12T16:39:21.133' AS DateTime))
GO
INSERT [dbo].[RoleMenus] ([role_id], [menu_id], [permission], [assigned_at]) VALUES (1, 1, N'view', CAST(N'2024-11-12T16:39:21.133' AS DateTime))
GO
INSERT [dbo].[RoleMenus] ([role_id], [menu_id], [permission], [assigned_at]) VALUES (1, 2, N'view', CAST(N'2024-11-12T16:39:21.133' AS DateTime))
GO
INSERT [dbo].[RoleMenus] ([role_id], [menu_id], [permission], [assigned_at]) VALUES (2, 3, N'edit', CAST(N'2024-11-12T16:39:21.133' AS DateTime))
GO
INSERT [dbo].[RoleMenus] ([role_id], [menu_id], [permission], [assigned_at]) VALUES (2, 3, N'view', CAST(N'2024-11-12T16:39:21.133' AS DateTime))
GO
INSERT [dbo].[RoleMenus] ([role_id], [menu_id], [permission], [assigned_at]) VALUES (3, 4, N'view', CAST(N'2024-11-12T16:39:21.133' AS DateTime))
GO
SET IDENTITY_INSERT [dbo].[Roles] ON 
GO
INSERT [dbo].[Roles] ([role_id], [name], [description], [created_at]) VALUES (1, N'Admin', N'Administrator with full access', CAST(N'2024-11-12T16:39:21.083' AS DateTime))
GO
INSERT [dbo].[Roles] ([role_id], [name], [description], [created_at]) VALUES (2, N'Editor', N'User with editing rights', CAST(N'2024-11-12T16:39:21.083' AS DateTime))
GO
INSERT [dbo].[Roles] ([role_id], [name], [description], [created_at]) VALUES (3, N'Viewer', N'User with view-only access', CAST(N'2024-11-12T16:39:21.083' AS DateTime))
GO
SET IDENTITY_INSERT [dbo].[Roles] OFF
GO
INSERT [dbo].[UserRoles] ([user_id], [role_id], [assigned_at]) VALUES (1, 1, CAST(N'2024-11-12T16:39:21.150' AS DateTime))
GO
INSERT [dbo].[UserRoles] ([user_id], [role_id], [assigned_at]) VALUES (2, 2, CAST(N'2024-11-12T16:39:21.150' AS DateTime))
GO
INSERT [dbo].[UserRoles] ([user_id], [role_id], [assigned_at]) VALUES (3, 3, CAST(N'2024-11-12T16:39:21.150' AS DateTime))
GO
SET IDENTITY_INSERT [dbo].[Users] ON 
GO
INSERT [dbo].[Users] ([user_id], [username], [email], [password], [status], [created_at], [updated_at]) VALUES (1, N'admin_user', N'admin@example.com', N'hashed_password_1', 1, CAST(N'2024-11-12T16:39:21.097' AS DateTime), NULL)
GO
INSERT [dbo].[Users] ([user_id], [username], [email], [password], [status], [created_at], [updated_at]) VALUES (2, N'editor_user', N'editor@example.com', N'hashed_password_2', 1, CAST(N'2024-11-12T16:39:21.097' AS DateTime), NULL)
GO
INSERT [dbo].[Users] ([user_id], [username], [email], [password], [status], [created_at], [updated_at]) VALUES (3, N'viewer_user', N'viewer@example.com', N'hashed_password_3', 1, CAST(N'2024-11-12T16:39:21.097' AS DateTime), NULL)
GO
INSERT [dbo].[Users] ([user_id], [username], [email], [password], [status], [created_at], [updated_at]) VALUES (4, N'nurulhadi', N'nurul.hadi@outlook.com', N'$2a$10$YLM17RU034Do91Cih0XRyu6.mrGQsMUc6Q1JgfJFsM/3il1HbEqNO', 1, CAST(N'2024-11-12T17:05:33.543' AS DateTime), NULL)
GO
SET IDENTITY_INSERT [dbo].[Users] OFF
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__Roles__72E12F1BC89F4B84]    Script Date: 13/11/2024 22:40:27 ******/
ALTER TABLE [dbo].[Roles] ADD UNIQUE NONCLUSTERED 
(
	[name] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__Users__AB6E6164023B840D]    Script Date: 13/11/2024 22:40:27 ******/
ALTER TABLE [dbo].[Users] ADD UNIQUE NONCLUSTERED 
(
	[email] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__Users__F3DBC572AB5C95CE]    Script Date: 13/11/2024 22:40:27 ******/
ALTER TABLE [dbo].[Users] ADD UNIQUE NONCLUSTERED 
(
	[username] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
ALTER TABLE [dbo].[Menus] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[RoleMenus] ADD  DEFAULT (getdate()) FOR [assigned_at]
GO
ALTER TABLE [dbo].[Roles] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[UserRoles] ADD  DEFAULT (getdate()) FOR [assigned_at]
GO
ALTER TABLE [dbo].[Users] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[Users] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[Menus]  WITH CHECK ADD  CONSTRAINT [FK_Menus_ParentID] FOREIGN KEY([parent_id])
REFERENCES [dbo].[Menus] ([menu_id])
GO
ALTER TABLE [dbo].[Menus] CHECK CONSTRAINT [FK_Menus_ParentID]
GO
ALTER TABLE [dbo].[RoleMenus]  WITH CHECK ADD  CONSTRAINT [FK_RoleMenus_MenuID] FOREIGN KEY([menu_id])
REFERENCES [dbo].[Menus] ([menu_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[RoleMenus] CHECK CONSTRAINT [FK_RoleMenus_MenuID]
GO
ALTER TABLE [dbo].[RoleMenus]  WITH CHECK ADD  CONSTRAINT [FK_RoleMenus_RoleID] FOREIGN KEY([role_id])
REFERENCES [dbo].[Roles] ([role_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[RoleMenus] CHECK CONSTRAINT [FK_RoleMenus_RoleID]
GO
ALTER TABLE [dbo].[UserRoles]  WITH CHECK ADD  CONSTRAINT [FK_UserRoles_RoleID] FOREIGN KEY([role_id])
REFERENCES [dbo].[Roles] ([role_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[UserRoles] CHECK CONSTRAINT [FK_UserRoles_RoleID]
GO
ALTER TABLE [dbo].[UserRoles]  WITH CHECK ADD  CONSTRAINT [FK_UserRoles_UserID] FOREIGN KEY([user_id])
REFERENCES [dbo].[Users] ([user_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[UserRoles] CHECK CONSTRAINT [FK_UserRoles_UserID]
GO
ALTER TABLE [dbo].[RoleMenus]  WITH CHECK ADD CHECK  (([permission]='delete' OR [permission]='edit' OR [permission]='view'))
GO
