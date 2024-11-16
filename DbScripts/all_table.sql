USE [AmanahDb]
GO
/****** Object:  Table [dbo].[BA]    Script Date: 14/11/2024 23:17:15 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[BA](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NULL,
	[ba_date] [date] NULL,
	[ba_subject] [varchar](255) NULL,
	[recepient_name] [varchar](255) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[BADetails]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[BADetails](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[section_id] [int] NULL,
	[item_name] [varchar](255) NULL,
	[quantity] [decimal](10, 2) NULL,
	[unit] [varchar](10) NULL,
	[weight_percentage] [decimal](5, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[BAProgress]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[BAProgress](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[detail_id] [int] NULL,
	[progress_previous_m2] [decimal](10, 2) NULL,
	[progress_previous_percentage] [decimal](5, 2) NULL,
	[progress_current_m2] [decimal](10, 2) NULL,
	[progress_current_percentage] [decimal](5, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[BASection]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[BASection](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[ba_id] [int] NULL,
	[section_name] [varchar](255) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[BreakdownItems]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[BreakdownItems](
	[item_id] [int] IDENTITY(1,1) NOT NULL,
	[section_id] [int] NULL,
	[description] [varchar](255) NULL,
	[unit_price] [decimal](15, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[item_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Breakdowns]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Breakdowns](
	[breakdown_id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NULL,
	[subject] [varchar](255) NULL,
	[location] [varchar](255) NULL,
	[date] [date] NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[breakdown_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[BreakdownSections]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[BreakdownSections](
	[section_id] [int] IDENTITY(1,1) NOT NULL,
	[breakdown_id] [int] NULL,
	[section_title] [varchar](255) NULL,
	[unit] [varchar](10) NULL,
	[total] [decimal](15, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[section_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Invoices]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Invoices](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[invoice_no] [varchar](50) NOT NULL,
	[invoice_date] [date] NOT NULL,
	[received_from] [varchar](100) NULL,
	[address] [text] NULL,
	[phone] [varchar](20) NULL,
	[fax] [varchar](20) NULL,
	[amount] [decimal](18, 2) NOT NULL,
	[payment_description] [varchar](255) NULL,
	[project_name] [varchar](255) NULL,
	[project_address] [text] NULL,
	[spk_no] [varchar](50) NULL,
	[contract_value] [decimal](18, 2) NULL,
	[progress_percentage] [decimal](5, 2) NULL,
	[progress_value] [decimal](18, 2) NULL,
	[bank_name] [varchar](50) NULL,
	[account_no] [varchar](50) NULL,
	[account_name] [varchar](100) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Menus]    Script Date: 14/11/2024 23:17:16 ******/
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
/****** Object:  Table [dbo].[OpnameDetail]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[OpnameDetail](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[section_id] [int] NOT NULL,
	[item_name] [varchar](255) NOT NULL,
	[quantity] [decimal](10, 2) NULL,
	[unit] [varchar](10) NULL,
	[unit_price] [decimal](18, 2) NULL,
	[total_price]  AS ([quantity]*[unit_price]) PERSISTED,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[OpnameMandor]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[OpnameMandor](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NOT NULL,
	[opname_name] [varchar](255) NOT NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[OpnamePayment]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[OpnamePayment](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[opname_mandor_id] [int] NOT NULL,
	[payment_description] [varchar](255) NULL,
	[payment_amount] [decimal](18, 2) NULL,
	[payment_date] [date] NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[OpnameRekap]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[OpnameRekap](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[opname_mandor_id] [int] NOT NULL,
	[total_payment] [decimal](18, 2) NULL,
	[remaining_payment] [decimal](18, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[OpnameSection]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[OpnameSection](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[opname_mandor_id] [int] NOT NULL,
	[section_name] [varchar](255) NOT NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[ProjectAdditionalExpenses]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[ProjectAdditionalExpenses](
	[expense_id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NULL,
	[expense_type] [varchar](50) NULL,
	[description] [varchar](255) NULL,
	[amount] [decimal](18, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[expense_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[ProjectCashFlow]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[ProjectCashFlow](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NOT NULL,
	[in_amount] [decimal](18, 2) NULL,
	[out_amount] [decimal](18, 2) NULL,
	[balance]  AS ([in_amount]-[out_amount]),
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[ProjectCashFlowDetails]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[ProjectCashFlowDetails](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[cashflow_id] [int] NOT NULL,
	[transaction_date] [date] NOT NULL,
	[description] [nvarchar](255) NULL,
	[in_amount] [decimal](18, 2) NULL,
	[out_amount] [decimal](18, 2) NULL,
	[balance] [decimal](18, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[ProjectRekap]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[ProjectRekap](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NOT NULL,
	[total_opname] [decimal](18, 2) NULL,
	[total_pengeluaran] [decimal](18, 2) NULL,
	[margin] [decimal](18, 2) NULL,
	[margin_percentage] [decimal](5, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Projects]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Projects](
	[project_id] [int] IDENTITY(1,1) NOT NULL,
	[project_name] [varchar](255) NOT NULL,
	[location] [varchar](255) NULL,
	[start_date] [date] NULL,
	[end_date] [date] NULL,
	[description] [text] NULL,
	[status] [varchar](20) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[project_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
/****** Object:  Table [dbo].[ProjectUser]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[ProjectUser](
	[id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NOT NULL,
	[user_id] [int] NOT NULL,
	[role] [nvarchar](50) NULL,
	[created_at] [datetime] NULL,
	[created_by] [int] NULL,
PRIMARY KEY CLUSTERED 
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[RekapPengeluaranDetails]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[RekapPengeluaranDetails](
	[detail_id] [int] IDENTITY(1,1) NOT NULL,
	[section_id] [int] NOT NULL,
	[item_description] [nvarchar](255) NOT NULL,
	[amount] [decimal](18, 2) NOT NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[detail_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[RekapPengeluaranSections]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[RekapPengeluaranSections](
	[section_id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NOT NULL,
	[section_name] [nvarchar](100) NOT NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[section_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[RoleMenus]    Script Date: 14/11/2024 23:17:16 ******/
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
/****** Object:  Table [dbo].[Roles]    Script Date: 14/11/2024 23:17:16 ******/
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
/****** Object:  Table [dbo].[Sph]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Sph](
	[sph_id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NULL,
	[recepient_name] [varchar](255) NULL,
	[date] [date] NULL,
	[subject] [varchar](255) NULL,
	[total_amount] [decimal](15, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[sph_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[SphDetails]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[SphDetails](
	[detail_id] [int] IDENTITY(1,1) NOT NULL,
	[section_id] [int] NULL,
	[item_description] [varchar](255) NULL,
	[quantity] [decimal](10, 2) NULL,
	[unit] [varchar](10) NULL,
	[unit_price] [decimal](15, 2) NULL,
	[discount_price] [decimal](15, 2) NULL,
	[total_price] [decimal](15, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[detail_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[SphSections]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[SphSections](
	[section_id] [int] IDENTITY(1,1) NOT NULL,
	[sph_id] [int] NULL,
	[section_title] [varchar](255) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[section_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Spk]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Spk](
	[spk_id] [int] IDENTITY(1,1) NOT NULL,
	[project_id] [int] NULL,
	[subject] [varchar](255) NULL,
	[date] [date] NULL,
	[total_jasa] [decimal](15, 2) NULL,
	[total_material] [decimal](15, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[spk_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[SPKDetails]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[SPKDetails](
	[detail_id] [int] IDENTITY(1,1) NOT NULL,
	[section_id] [int] NULL,
	[description] [varchar](255) NULL,
	[quantity] [decimal](10, 2) NULL,
	[unit] [varchar](10) NULL,
	[unit_price_jasa] [decimal](15, 2) NULL,
	[total_jasa] [decimal](15, 2) NULL,
	[unit_price_material] [decimal](15, 2) NULL,
	[total_material] [decimal](15, 2) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[detail_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[SPKSections]    Script Date: 14/11/2024 23:17:16 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[SPKSections](
	[section_id] [int] IDENTITY(1,1) NOT NULL,
	[spk_id] [int] NULL,
	[section_title] [varchar](255) NULL,
	[created_by] [int] NULL,
	[created_at] [datetime] NULL,
	[updated_by] [int] NULL,
	[updated_at] [datetime] NULL,
	[deleted_by] [int] NULL,
	[deleted_at] [datetime] NULL,
PRIMARY KEY CLUSTERED 
(
	[section_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[UserRoles]    Script Date: 14/11/2024 23:17:16 ******/
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
/****** Object:  Table [dbo].[Users]    Script Date: 14/11/2024 23:17:16 ******/
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
SET IDENTITY_INSERT [dbo].[BA] ON 
GO
INSERT [dbo].[BA] ([id], [project_id], [ba_date], [ba_subject], [recepient_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, CAST(N'2022-04-11' AS Date), N'Berita Acara Progres pekerjaan 1', N'Bp Randy Andrian', 1, CAST(N'2024-11-14T15:17:51.870' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[BA] OFF
GO
SET IDENTITY_INSERT [dbo].[BADetails] ON 
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Jasa pasang lantai', CAST(16.00 AS Decimal(10, 2)), N'm2', CAST(5.43 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Jasa pasang dinding', CAST(42.76 AS Decimal(10, 2)), N'm2', CAST(14.79 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'Jasa pasang toping / ambalan', CAST(9.20 AS Decimal(10, 2)), N'ml', CAST(2.12 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, N'Jasa pasang top vanity & pelubangan', CAST(4.60 AS Decimal(10, 2)), N'ml', CAST(2.12 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, N'Jasa pasang & material rangka boxes vanity', CAST(5.00 AS Decimal(10, 2)), N'unit', CAST(3.45 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (6, 2, N'Jasa pasang lantai', CAST(21.02 AS Decimal(10, 2)), N'm2', CAST(7.14 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (7, 2, N'Jasa pasang dinding', CAST(61.08 AS Decimal(10, 2)), N'm2', CAST(21.12 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (8, 2, N'Jasa pasang toping / ambalan', CAST(7.40 AS Decimal(10, 2)), N'ml', CAST(1.70 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (9, 2, N'Jasa pasang top vanity & pelubangan', CAST(3.85 AS Decimal(10, 2)), N'ml', CAST(1.77 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (10, 2, N'Jasa pasang & material rangka boxes vanity', CAST(4.00 AS Decimal(10, 2)), N'unit', CAST(2.76 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:20:53.640' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (11, 1, N'Jasa pasang lantai', CAST(16.00 AS Decimal(10, 2)), N'm2', CAST(5.43 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (12, 1, N'Jasa pasang dinding', CAST(42.76 AS Decimal(10, 2)), N'm2', CAST(14.79 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (13, 1, N'Jasa pasang toping / ambalan', CAST(9.20 AS Decimal(10, 2)), N'm1', CAST(2.12 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (14, 1, N'Jasa pasang top vanity & pelubangan', CAST(4.60 AS Decimal(10, 2)), N'm1', CAST(2.12 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (15, 1, N'Jasa pasang & material rangka boxes vanity', CAST(5.00 AS Decimal(10, 2)), N'unit', CAST(3.45 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (16, 2, N'Jasa pasang lantai', CAST(21.02 AS Decimal(10, 2)), N'm2', CAST(7.14 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (17, 2, N'Jasa pasang dinding', CAST(61.08 AS Decimal(10, 2)), N'm2', CAST(21.12 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (18, 2, N'Jasa pasang toping / ambalan', CAST(7.40 AS Decimal(10, 2)), N'm1', CAST(1.70 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (19, 2, N'Jasa pasang top vanity & pelubangan', CAST(3.85 AS Decimal(10, 2)), N'm1', CAST(1.77 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (20, 2, N'Jasa pasang & material rangka boxes vanity', CAST(4.00 AS Decimal(10, 2)), N'unit', CAST(2.76 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (21, 3, N'Jasa pasang lantai', CAST(4.02 AS Decimal(10, 2)), N'm2', CAST(1.36 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (22, 3, N'Jasa pasang dinding', CAST(23.25 AS Decimal(10, 2)), N'm2', CAST(8.04 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (23, 3, N'Jasa pasang toping / ambalan', CAST(7.44 AS Decimal(10, 2)), N'm1', CAST(1.71 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (24, 3, N'Jasa pasang top vanity & pelubangan', CAST(0.65 AS Decimal(10, 2)), N'm1', CAST(0.30 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (25, 3, N'Penebalan rangka dinding (dry sistem)', CAST(5.58 AS Decimal(10, 2)), N'm2', CAST(2.10 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (26, 4, N'Jasa pasang lantai', CAST(4.68 AS Decimal(10, 2)), N'm2', CAST(1.59 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (27, 4, N'Jasa pasang dinding', CAST(23.94 AS Decimal(10, 2)), N'm2', CAST(8.28 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (28, 4, N'Jasa pasang toping / ambalan', CAST(7.45 AS Decimal(10, 2)), N'm1', CAST(1.72 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (29, 4, N'Jasa pasang top vanity & pelubangan', CAST(0.65 AS Decimal(10, 2)), N'm1', CAST(0.30 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (30, 4, N'Jasa pasang & material rangka boxes vanity', CAST(1.00 AS Decimal(10, 2)), N'unit', CAST(0.69 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (31, 5, N'Mobilisasi, kebersihan dll', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(3.45 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (32, 5, N'Persiapan & kerja peralatan', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(1.08 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BADetails] ([id], [section_id], [item_name], [quantity], [unit], [weight_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (33, 5, N'Koordinasi & supervisi', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(0.69 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:29:55.800' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[BADetails] OFF
GO
SET IDENTITY_INSERT [dbo].[BAProgress] ON 
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(13.62 AS Decimal(10, 2)), CAST(4.62 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 2, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(32.42 AS Decimal(10, 2)), CAST(11.21 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 3, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(4.58 AS Decimal(10, 2)), CAST(1.05 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 4, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 5, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (6, 6, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(12.92 AS Decimal(10, 2)), CAST(4.39 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (7, 7, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(61.31 AS Decimal(10, 2)), CAST(21.20 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (8, 8, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(12.04 AS Decimal(10, 2)), CAST(2.77 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (9, 9, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (10, 10, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (11, 11, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (12, 12, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(16.15 AS Decimal(10, 2)), CAST(5.59 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (13, 13, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(5.21 AS Decimal(10, 2)), CAST(1.20 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (14, 14, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (15, 15, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(6.82 AS Decimal(10, 2)), CAST(2.56 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (16, 16, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(4.47 AS Decimal(10, 2)), CAST(1.52 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (17, 17, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(23.73 AS Decimal(10, 2)), CAST(8.21 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (18, 18, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(7.43 AS Decimal(10, 2)), CAST(1.71 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (19, 19, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (20, 20, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (21, 21, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (22, 22, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BAProgress] ([id], [detail_id], [progress_previous_m2], [progress_previous_percentage], [progress_current_m2], [progress_current_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (23, 23, CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), CAST(0.00 AS Decimal(10, 2)), CAST(0.00 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:33:09.523' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[BAProgress] OFF
GO
SET IDENTITY_INSERT [dbo].[BASection] ON 
GO
INSERT [dbo].[BASection] ([id], [ba_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Female Toilet', 1, CAST(N'2024-11-14T15:19:36.660' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BASection] ([id], [ba_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Male Toilet', 1, CAST(N'2024-11-14T15:19:36.660' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BASection] ([id], [ba_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'Difable Toilet', 1, CAST(N'2024-11-14T15:19:36.660' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BASection] ([id], [ba_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, N'Nursery', 1, CAST(N'2024-11-14T15:19:36.660' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BASection] ([id], [ba_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, N'Preliminaries', 1, CAST(N'2024-11-14T15:19:36.660' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[BASection] OFF
GO
SET IDENTITY_INSERT [dbo].[BreakdownItems] ON 
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Jasa setting & Pasang Material', CAST(243000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'screed mu 301 s3 / setara (tebal 30mm)', CAST(97050.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'Tile Adhesive MU 450 / setara – Tebal max 5 mm', CAST(47500.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, N'Perekat Marmer Ke Adakan Latricrete L3642 (bonding agent)', CAST(85000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, N'Poles Marmer fin. polish + Resin', CAST(115000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (6, 1, N'Epoxy / RTN', CAST(45000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (7, 1, N'Proteksi plastik cor', CAST(11000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (8, 1, N'Angkut', CAST(33000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (9, 1, N'Coating Material eks Akemi / FILA / shinol', CAST(67500.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.907' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (10, 2, N'Jasa setting & Pasang Material', CAST(735000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.910' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (11, 2, N'Perekat marmer', CAST(24000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.910' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (12, 2, N'detail edging minimalis', CAST(75000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.910' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (13, 2, N'Angkut', CAST(87500.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.910' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (14, 2, N'Coating Material eks Akemi / ishinol', CAST(67500.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.910' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (15, 3, N'pekerjaan detail profile', CAST(575000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.913' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (16, 3, N'pekerjaan detail bullnose', CAST(105000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.913' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (17, 3, N'pekerjaan detail tali air', CAST(162000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.913' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (18, 3, N'pekerjaan detail outter', CAST(192000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.913' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (19, 3, N'pekerjaan detail inlay', CAST(360000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.913' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (20, 4, N'Mobilisasi, kebersihan dll', CAST(3540000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.920' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (21, 4, N'Supervisi, gambar dan pemilihan marmer', CAST(1200000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.920' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownItems] ([item_id], [section_id], [description], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (22, 4, N'K3 dan Id Card', CAST(165000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:35:34.920' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[BreakdownItems] OFF
GO
SET IDENTITY_INSERT [dbo].[Breakdowns] ON 
GO
INSERT [dbo].[Breakdowns] ([breakdown_id], [project_id], [subject], [location], [date], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Breakdown Pemasangan Marmer & Bahan Bantu', N'Jabodetabek', CAST(N'2021-03-27' AS Date), 1, CAST(N'2024-11-14T14:34:28.737' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[Breakdowns] OFF
GO
SET IDENTITY_INSERT [dbo].[BreakdownSections] ON 
GO
INSERT [dbo].[BreakdownSections] ([section_id], [breakdown_id], [section_title], [unit], [total], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'PERINCIAN PEMASANGAN LANTAI', N'M2', CAST(677550.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:34:41.390' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownSections] ([section_id], [breakdown_id], [section_title], [unit], [total], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'PERINCIAN PEMASANGAN COUNTER TOP / VANITY TOP', N'M1', CAST(919500.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:34:41.397' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownSections] ([section_id], [breakdown_id], [section_title], [unit], [total], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'PERINCIAN PEKERJAAN DETAIL', N'M1', CAST(575000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:34:41.400' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownSections] ([section_id], [breakdown_id], [section_title], [unit], [total], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, N'PRELIMINARIES', N'LS', CAST(2929000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:34:41.403' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownSections] ([section_id], [breakdown_id], [section_title], [unit], [total], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, N'PERINCIAN PEMASANGAN LANTAI', N'M2', CAST(677550.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:34:41.407' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownSections] ([section_id], [breakdown_id], [section_title], [unit], [total], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (6, 1, N'PERINCIAN PEMASANGAN COUNTER TOP / VANITY TOP', N'M1', CAST(919500.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:34:41.413' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownSections] ([section_id], [breakdown_id], [section_title], [unit], [total], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (7, 1, N'PERINCIAN PEKERJAAN DETAIL', N'M1', CAST(575000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:34:41.413' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[BreakdownSections] ([section_id], [breakdown_id], [section_title], [unit], [total], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (8, 1, N'PRELIMINARIES', N'LS', CAST(2929000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:34:41.417' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[BreakdownSections] OFF
GO
SET IDENTITY_INSERT [dbo].[Invoices] ON 
GO
INSERT [dbo].[Invoices] ([id], [invoice_no], [invoice_date], [received_from], [address], [phone], [fax], [amount], [payment_description], [project_name], [project_address], [spk_no], [contract_value], [progress_percentage], [progress_value], [bank_name], [account_no], [account_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, N'Kw/II/ASA/356/2022', CAST(N'2022-04-04' AS Date), N'PT. Wijaya Kusuma Contractors', N'Jl. R.P. Soeroso No. 32, Jakarta', N'0213905658', N'', CAST(95430392.00 AS Decimal(18, 2)), N'progres 1 Pekerjaan pasang marmer toilet magran - T2', N'Indonesia Design District, Jl Rasuna Said, Distrik 11PIK - 2', N'Jl Rasuna Said, Jakarta', N'NO .../SPK-SATINO/WKC-ID/IV/2023', CAST(175650638.00 AS Decimal(18, 2)), CAST(54.33 AS Decimal(5, 2)), CAST(95430392.00 AS Decimal(18, 2)), N'BCA', N'3720501111', N'Pilar Asa Mandiri PT', 1, CAST(N'2024-11-14T15:37:29.617' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[Invoices] OFF
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
SET IDENTITY_INSERT [dbo].[OpnameDetail] ON 
GO
INSERT [dbo].[OpnameDetail] ([id], [section_id], [item_name], [quantity], [unit], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Jasa pasang lantai', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(4500000.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:43:39.303' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[OpnameDetail] ([id], [section_id], [item_name], [quantity], [unit], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Jasa pasang dinding', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(4500000.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:43:39.303' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[OpnameDetail] ([id], [section_id], [item_name], [quantity], [unit], [unit_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 2, N'Pekerjaan detail nursery', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(2000000.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:43:39.303' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[OpnameDetail] OFF
GO
SET IDENTITY_INSERT [dbo].[OpnameMandor] ON 
GO
INSERT [dbo].[OpnameMandor] ([id], [project_id], [opname_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'OPNAME PEKERJAAN MANDOR RUSKI', 1, CAST(N'2024-11-14T15:43:17.173' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[OpnameMandor] OFF
GO
SET IDENTITY_INSERT [dbo].[OpnamePayment] ON 
GO
INSERT [dbo].[OpnamePayment] ([id], [opname_mandor_id], [payment_description], [payment_amount], [payment_date], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Pembayaran 1', CAST(5000000.00 AS Decimal(18, 2)), CAST(N'2022-01-01' AS Date), 1, CAST(N'2024-11-14T15:44:36.443' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[OpnamePayment] ([id], [opname_mandor_id], [payment_description], [payment_amount], [payment_date], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Pembayaran 2', CAST(2000000.00 AS Decimal(18, 2)), CAST(N'2022-02-01' AS Date), 1, CAST(N'2024-11-14T15:44:36.443' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[OpnamePayment] OFF
GO
SET IDENTITY_INSERT [dbo].[OpnameRekap] ON 
GO
INSERT [dbo].[OpnameRekap] ([id], [opname_mandor_id], [total_payment], [remaining_payment], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, CAST(7000000.00 AS Decimal(18, 2)), CAST(7000000.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:45:03.163' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[OpnameRekap] OFF
GO
SET IDENTITY_INSERT [dbo].[OpnameSection] ON 
GO
INSERT [dbo].[OpnameSection] ([id], [opname_mandor_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Toilet', 1, CAST(N'2024-11-14T15:43:34.923' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[OpnameSection] ([id], [opname_mandor_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Nursery', 1, CAST(N'2024-11-14T15:43:34.923' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[OpnameSection] OFF
GO
SET IDENTITY_INSERT [dbo].[ProjectAdditionalExpenses] ON 
GO
INSERT [dbo].[ProjectAdditionalExpenses] ([expense_id], [project_id], [expense_type], [description], [amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Bahan Bantu & OPR', N'Biaya bahan bantu dan operasional', CAST(37015200.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T14:56:54.110' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectAdditionalExpenses] ([expense_id], [project_id], [expense_type], [description], [amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Lain-lain', N'Biaya tambahan lainnya', CAST(10000000.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T14:56:54.110' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[ProjectAdditionalExpenses] OFF
GO
SET IDENTITY_INSERT [dbo].[ProjectCashFlow] ON 
GO
INSERT [dbo].[ProjectCashFlow] ([id], [project_id], [in_amount], [out_amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, CAST(30000000.00 AS Decimal(18, 2)), CAST(34660290.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:58:27.140' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[ProjectCashFlow] OFF
GO
SET IDENTITY_INSERT [dbo].[ProjectCashFlowDetails] ON 
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, CAST(N'2022-09-26' AS Date), N'Pemasukan Awal', CAST(2000000.00 AS Decimal(18, 2)), CAST(0.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, CAST(N'2022-10-09' AS Date), N'Jaring paranet', CAST(0.00 AS Decimal(18, 2)), CAST(647300.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, CAST(N'2022-10-12' AS Date), N'Kopi kenangan', CAST(0.00 AS Decimal(18, 2)), CAST(159000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, CAST(N'2022-10-21' AS Date), N'BBM', CAST(0.00 AS Decimal(18, 2)), CAST(46000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, CAST(N'2022-10-25' AS Date), N'Print + BBM', CAST(0.00 AS Decimal(18, 2)), CAST(50000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (6, 1, CAST(N'2022-10-28' AS Date), N'BBM', CAST(0.00 AS Decimal(18, 2)), CAST(44000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (7, 1, CAST(N'2022-11-02' AS Date), N'BBM', CAST(0.00 AS Decimal(18, 2)), CAST(40000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (8, 1, CAST(N'2022-11-04' AS Date), N'BBM', CAST(0.00 AS Decimal(18, 2)), CAST(40000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (9, 1, CAST(N'2022-11-23' AS Date), N'BBM', CAST(0.00 AS Decimal(18, 2)), CAST(53000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (10, 1, CAST(N'2022-12-02' AS Date), N'Intertain', CAST(0.00 AS Decimal(18, 2)), CAST(60000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (11, 1, CAST(N'2023-01-01' AS Date), N'BBM', CAST(0.00 AS Decimal(18, 2)), CAST(48000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (12, 1, CAST(N'2023-02-09' AS Date), N'BBM', CAST(0.00 AS Decimal(18, 2)), CAST(30000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (13, 1, CAST(N'2023-02-13' AS Date), N'BBM', CAST(0.00 AS Decimal(18, 2)), CAST(50000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (14, 1, CAST(N'2024-05-10' AS Date), N'BB ganti nat resin master', CAST(0.00 AS Decimal(18, 2)), CAST(500000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (15, 1, CAST(N'2024-10-12' AS Date), N'Intertain', CAST(0.00 AS Decimal(18, 2)), CAST(50000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[ProjectCashFlowDetails] ([id], [cashflow_id], [transaction_date], [description], [in_amount], [out_amount], [balance], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (16, 1, CAST(N'2024-10-26' AS Date), N'Byr tukang & alat treatment flek', CAST(0.00 AS Decimal(18, 2)), CAST(1227000.00 AS Decimal(18, 2)), NULL, 1, CAST(N'2024-11-14T16:02:38.887' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[ProjectCashFlowDetails] OFF
GO
SET IDENTITY_INSERT [dbo].[ProjectRekap] ON 
GO
INSERT [dbo].[ProjectRekap] ([id], [project_id], [total_opname], [total_pengeluaran], [margin], [margin_percentage], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, CAST(175650638.00 AS Decimal(18, 2)), CAST(102749850.00 AS Decimal(18, 2)), CAST(45677670.00 AS Decimal(18, 2)), CAST(30.77 AS Decimal(5, 2)), 1, CAST(N'2024-11-14T15:54:32.397' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[ProjectRekap] OFF
GO
SET IDENTITY_INSERT [dbo].[Projects] ON 
GO
INSERT [dbo].[Projects] ([project_id], [project_name], [location], [start_date], [end_date], [description], [status], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, N'Proyek Gedung A', N'Jakarta, Indonesia', CAST(N'2024-01-01' AS Date), CAST(N'2024-12-31' AS Date), N'Pembangunan gedung perkantoran', N'in-progress', 1, CAST(N'2024-11-14T14:27:00.500' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[Projects] ([project_id], [project_name], [location], [start_date], [end_date], [description], [status], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, N'Proyek Mall B', N'Bandung, Indonesia', CAST(N'2024-03-01' AS Date), CAST(N'2024-11-30' AS Date), N'Renovasi pusat perbelanjaan', N'in-progress', 2, CAST(N'2024-11-14T14:27:00.500' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[Projects] OFF
GO
SET IDENTITY_INSERT [dbo].[RekapPengeluaranDetails] ON 
GO
INSERT [dbo].[RekapPengeluaranDetails] ([detail_id], [section_id], [item_description], [amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'mandor ipung', CAST(30017500.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:54:56.100' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[RekapPengeluaranDetails] ([detail_id], [section_id], [item_description], [amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'mandor musholi', CAST(38326100.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:54:56.100' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[RekapPengeluaranDetails] ([detail_id], [section_id], [item_description], [amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'mandor ruski', CAST(14000000.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:54:56.100' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[RekapPengeluaranDetails] ([detail_id], [section_id], [item_description], [amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 2, N'Bahan Bantu', CAST(20406250.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:54:56.100' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[RekapPengeluaranDetails] ([detail_id], [section_id], [item_description], [amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 3, N'Lain-lain', CAST(10000000.00 AS Decimal(18, 2)), 1, CAST(N'2024-11-14T15:54:56.100' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[RekapPengeluaranDetails] OFF
GO
SET IDENTITY_INSERT [dbo].[RekapPengeluaranSections] ON 
GO
INSERT [dbo].[RekapPengeluaranSections] ([section_id], [project_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Opname Mandor', 1, CAST(N'2024-11-14T15:54:44.157' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[RekapPengeluaranSections] ([section_id], [project_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Bahan Bantu & OPR', 1, CAST(N'2024-11-14T15:54:44.157' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[RekapPengeluaranSections] ([section_id], [project_id], [section_name], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'Lain-lain', 1, CAST(N'2024-11-14T15:54:44.157' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[RekapPengeluaranSections] OFF
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
SET IDENTITY_INSERT [dbo].[Sph] ON 
GO
INSERT [dbo].[Sph] ([sph_id], [project_id], [recepient_name], [date], [subject], [total_amount], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Bp Randy Andrian', CAST(N'2021-03-27' AS Date), N'Penawaran Harga rev 3', CAST(175550630.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:47:45.723' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[Sph] OFF
GO
SET IDENTITY_INSERT [dbo].[SphDetails] ON 
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Jasa pasang lantai', CAST(16.00 AS Decimal(10, 2)), N'm2', CAST(677550.00 AS Decimal(15, 2)), CAST(596244.00 AS Decimal(15, 2)), CAST(9539904.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:31.797' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Jasa pasang dinding', CAST(42.76 AS Decimal(10, 2)), N'm2', CAST(690300.00 AS Decimal(15, 2)), CAST(607464.00 AS Decimal(15, 2)), CAST(25975161.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:31.797' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'Jasa pasang toping / ambalan', CAST(9.20 AS Decimal(10, 2)), N'm1', CAST(459750.00 AS Decimal(15, 2)), CAST(404580.00 AS Decimal(15, 2)), CAST(3722136.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:31.797' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, N'Jasa pasang top vanity & pelubangan', CAST(4.60 AS Decimal(10, 2)), N'm1', CAST(919500.00 AS Decimal(15, 2)), CAST(809160.00 AS Decimal(15, 2)), CAST(3722136.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:31.797' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, N'Jasa pasang & material rangka boxes vanity', CAST(5.00 AS Decimal(10, 2)), N'unit', CAST(1379250.00 AS Decimal(15, 2)), CAST(1213740.00 AS Decimal(15, 2)), CAST(6068700.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:31.797' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (6, 2, N'Jasa pasang lantai', CAST(21.02 AS Decimal(10, 2)), N'm2', CAST(677550.00 AS Decimal(15, 2)), CAST(596244.00 AS Decimal(15, 2)), CAST(12533049.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:47.603' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (7, 2, N'Jasa pasang dinding', CAST(61.08 AS Decimal(10, 2)), N'm2', CAST(690300.00 AS Decimal(15, 2)), CAST(607464.00 AS Decimal(15, 2)), CAST(37103901.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:47.603' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (8, 2, N'Jasa pasang toping / ambalan', CAST(7.40 AS Decimal(10, 2)), N'm1', CAST(459750.00 AS Decimal(15, 2)), CAST(404580.00 AS Decimal(15, 2)), CAST(2933832.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:47.603' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (9, 2, N'Jasa pasang top vanity & pelubangan', CAST(3.85 AS Decimal(10, 2)), N'm1', CAST(919500.00 AS Decimal(15, 2)), CAST(809160.00 AS Decimal(15, 2)), CAST(3115266.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:47.603' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (10, 2, N'Jasa pasang & material rangka boxes vanity', CAST(4.00 AS Decimal(10, 2)), N'unit', CAST(1379250.00 AS Decimal(15, 2)), CAST(1213740.00 AS Decimal(15, 2)), CAST(4854960.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:47.603' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (11, 3, N'Jasa pasang lantai', CAST(4.02 AS Decimal(10, 2)), N'm2', CAST(677550.00 AS Decimal(15, 2)), CAST(596244.00 AS Decimal(15, 2)), CAST(2396901.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:53.277' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (12, 3, N'Jasa pasang dinding', CAST(23.25 AS Decimal(10, 2)), N'm2', CAST(690300.00 AS Decimal(15, 2)), CAST(607464.00 AS Decimal(15, 2)), CAST(14123538.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:53.277' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (13, 3, N'Jasa pasang toping / ambalan', CAST(7.44 AS Decimal(10, 2)), N'm1', CAST(459750.00 AS Decimal(15, 2)), CAST(404580.00 AS Decimal(15, 2)), CAST(3255954.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:53.277' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (14, 3, N'Jasa pasang top vanity & pelubangan', CAST(0.65 AS Decimal(10, 2)), N'm1', CAST(919500.00 AS Decimal(15, 2)), CAST(809160.00 AS Decimal(15, 2)), CAST(525954.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:53.277' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (15, 3, N'Penebalan rangka dinding (dry sistem)', CAST(5.58 AS Decimal(10, 2)), N'm2', CAST(750000.00 AS Decimal(15, 2)), CAST(660000.00 AS Decimal(15, 2)), CAST(3682800.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:48:53.277' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (16, 4, N'Jasa pasang lantai', CAST(4.68 AS Decimal(10, 2)), N'm2', CAST(677550.00 AS Decimal(15, 2)), CAST(596244.00 AS Decimal(15, 2)), CAST(2790402.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:50:10.813' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (17, 4, N'Jasa pasang dinding', CAST(23.94 AS Decimal(10, 2)), N'm2', CAST(690300.00 AS Decimal(15, 2)), CAST(607464.00 AS Decimal(15, 2)), CAST(14542688.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:50:10.813' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (18, 4, N'Jasa pasang toping / ambalan', CAST(7.45 AS Decimal(10, 2)), N'm1', CAST(459750.00 AS Decimal(15, 2)), CAST(404580.00 AS Decimal(15, 2)), CAST(3014121.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:50:10.813' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (19, 4, N'Jasa pasang top vanity & pelubangan', CAST(0.65 AS Decimal(10, 2)), N'm1', CAST(919500.00 AS Decimal(15, 2)), CAST(809160.00 AS Decimal(15, 2)), CAST(525954.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:50:10.813' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (20, 4, N'Jasa pasang & material rangka boxes vanity', CAST(1.00 AS Decimal(10, 2)), N'unit', CAST(1379250.00 AS Decimal(15, 2)), CAST(1213740.00 AS Decimal(15, 2)), CAST(1213740.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:50:10.813' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (21, 5, N'Mobilisasi, kebersihan dll', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(3870000.00 AS Decimal(15, 2)), CAST(3405600.00 AS Decimal(15, 2)), CAST(3405600.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:50:42.970' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (22, 5, N'Persiapan & kerja peralatan', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(1200000.00 AS Decimal(15, 2)), CAST(1056000.00 AS Decimal(15, 2)), CAST(1056000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:50:42.970' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphDetails] ([detail_id], [section_id], [item_description], [quantity], [unit], [unit_price], [discount_price], [total_price], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (23, 5, N'Koordinasi & supervisi', CAST(1.00 AS Decimal(10, 2)), N'ls', CAST(16500000.00 AS Decimal(15, 2)), CAST(14520000.00 AS Decimal(15, 2)), CAST(14520000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:50:42.970' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[SphDetails] OFF
GO
SET IDENTITY_INSERT [dbo].[SphSections] ON 
GO
INSERT [dbo].[SphSections] ([section_id], [sph_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Female Toilet', 1, CAST(N'2024-11-14T14:48:22.747' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphSections] ([section_id], [sph_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Male Toilet', 1, CAST(N'2024-11-14T14:48:22.753' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphSections] ([section_id], [sph_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'Difable Toilet', 1, CAST(N'2024-11-14T14:48:22.757' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphSections] ([section_id], [sph_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, N'Nursery', 1, CAST(N'2024-11-14T14:48:22.763' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SphSections] ([section_id], [sph_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, N'Preliminaries', 1, CAST(N'2024-11-14T14:48:22.767' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[SphSections] OFF
GO
SET IDENTITY_INSERT [dbo].[Spk] ON 
GO
INSERT [dbo].[Spk] ([spk_id], [project_id], [subject], [date], [total_jasa], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'SPK Mandor', CAST(N'2021-03-27' AS Date), CAST(72274300.00 AS Decimal(15, 2)), CAST(37015200.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:39:50.903' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[Spk] OFF
GO
SET IDENTITY_INSERT [dbo].[SPKDetails] ON 
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Jasa pasang lantai', CAST(16.00 AS Decimal(10, 2)), N'm2', CAST(170000.00 AS Decimal(15, 2)), CAST(2720000.00 AS Decimal(15, 2)), CAST(100000.00 AS Decimal(15, 2)), CAST(1600000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.990' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Jasa pasang dinding', CAST(42.76 AS Decimal(10, 2)), N'm2', CAST(250000.00 AS Decimal(15, 2)), CAST(10690000.00 AS Decimal(15, 2)), CAST(140000.00 AS Decimal(15, 2)), CAST(5986400.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.990' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'Jasa pasang toping / ambalan', CAST(9.20 AS Decimal(10, 2)), N'ml', CAST(250000.00 AS Decimal(15, 2)), CAST(2300000.00 AS Decimal(15, 2)), CAST(100000.00 AS Decimal(15, 2)), CAST(920000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.990' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, N'Jasa pasang top vanity & pelubangan', CAST(4.60 AS Decimal(10, 2)), N'ml', CAST(450000.00 AS Decimal(15, 2)), CAST(2070000.00 AS Decimal(15, 2)), CAST(200000.00 AS Decimal(15, 2)), CAST(920000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.990' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, N'Jasa pasang & material rangka boxes vanity', CAST(5.00 AS Decimal(10, 2)), N'unit', CAST(400000.00 AS Decimal(15, 2)), CAST(2000000.00 AS Decimal(15, 2)), CAST(200000.00 AS Decimal(15, 2)), CAST(1000000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.990' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (6, 2, N'Jasa pasang lantai', CAST(21.02 AS Decimal(10, 2)), N'm2', CAST(260000.00 AS Decimal(15, 2)), CAST(5465200.00 AS Decimal(15, 2)), CAST(100000.00 AS Decimal(15, 2)), CAST(2102000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.993' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (7, 2, N'Jasa pasang dinding', CAST(61.08 AS Decimal(10, 2)), N'm2', CAST(330000.00 AS Decimal(15, 2)), CAST(20156400.00 AS Decimal(15, 2)), CAST(140000.00 AS Decimal(15, 2)), CAST(8551200.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.993' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (8, 2, N'Jasa pasang toping / ambalan', CAST(7.40 AS Decimal(10, 2)), N'ml', CAST(90000.00 AS Decimal(15, 2)), CAST(666000.00 AS Decimal(15, 2)), CAST(100000.00 AS Decimal(15, 2)), CAST(740000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.993' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (9, 2, N'Jasa pasang top vanity & pelubangan', CAST(3.85 AS Decimal(10, 2)), N'ml', CAST(450000.00 AS Decimal(15, 2)), CAST(1732500.00 AS Decimal(15, 2)), CAST(200000.00 AS Decimal(15, 2)), CAST(770000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.993' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKDetails] ([detail_id], [section_id], [description], [quantity], [unit], [unit_price_jasa], [total_jasa], [unit_price_material], [total_material], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (10, 2, N'Jasa pasang & material rangka boxes vanity', CAST(4.00 AS Decimal(10, 2)), N'unit', CAST(400000.00 AS Decimal(15, 2)), CAST(1600000.00 AS Decimal(15, 2)), CAST(200000.00 AS Decimal(15, 2)), CAST(800000.00 AS Decimal(15, 2)), 1, CAST(N'2024-11-14T14:40:05.993' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[SPKDetails] OFF
GO
SET IDENTITY_INSERT [dbo].[SPKSections] ON 
GO
INSERT [dbo].[SPKSections] ([section_id], [spk_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (1, 1, N'Female Toilet', 1, CAST(N'2024-11-14T14:39:57.237' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKSections] ([section_id], [spk_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (2, 1, N'Male Toilet', 1, CAST(N'2024-11-14T14:39:57.247' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKSections] ([section_id], [spk_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (3, 1, N'Difable Toilet', 1, CAST(N'2024-11-14T14:39:57.250' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKSections] ([section_id], [spk_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (4, 1, N'Nursery', 1, CAST(N'2024-11-14T14:39:57.257' AS DateTime), NULL, NULL, NULL, NULL)
GO
INSERT [dbo].[SPKSections] ([section_id], [spk_id], [section_title], [created_by], [created_at], [updated_by], [updated_at], [deleted_by], [deleted_at]) VALUES (5, 1, N'Preliminaries', 1, CAST(N'2024-11-14T14:39:57.260' AS DateTime), NULL, NULL, NULL, NULL)
GO
SET IDENTITY_INSERT [dbo].[SPKSections] OFF
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
/****** Object:  Index [UQ__Roles__72E12F1BC89F4B84]    Script Date: 14/11/2024 23:17:16 ******/
ALTER TABLE [dbo].[Roles] ADD UNIQUE NONCLUSTERED 
(
	[name] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__Users__AB6E6164023B840D]    Script Date: 14/11/2024 23:17:16 ******/
ALTER TABLE [dbo].[Users] ADD UNIQUE NONCLUSTERED 
(
	[email] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
/****** Object:  Index [UQ__Users__F3DBC572AB5C95CE]    Script Date: 14/11/2024 23:17:16 ******/
ALTER TABLE [dbo].[Users] ADD UNIQUE NONCLUSTERED 
(
	[username] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
GO
ALTER TABLE [dbo].[BA] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[BADetails] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[BAProgress] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[BASection] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[BreakdownItems] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[Breakdowns] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[BreakdownSections] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[Invoices] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[Menus] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[ProjectAdditionalExpenses] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[ProjectCashFlow] ADD  DEFAULT ((0)) FOR [in_amount]
GO
ALTER TABLE [dbo].[ProjectCashFlow] ADD  DEFAULT ((0)) FOR [out_amount]
GO
ALTER TABLE [dbo].[ProjectCashFlow] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[ProjectCashFlowDetails] ADD  DEFAULT ((0)) FOR [in_amount]
GO
ALTER TABLE [dbo].[ProjectCashFlowDetails] ADD  DEFAULT ((0)) FOR [out_amount]
GO
ALTER TABLE [dbo].[ProjectCashFlowDetails] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[Projects] ADD  DEFAULT ('in-progress') FOR [status]
GO
ALTER TABLE [dbo].[Projects] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[ProjectUser] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[RoleMenus] ADD  DEFAULT (getdate()) FOR [assigned_at]
GO
ALTER TABLE [dbo].[Roles] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[Sph] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[SphDetails] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[SphSections] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[Spk] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[SPKDetails] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[SPKSections] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[UserRoles] ADD  DEFAULT (getdate()) FOR [assigned_at]
GO
ALTER TABLE [dbo].[Users] ADD  DEFAULT ((1)) FOR [status]
GO
ALTER TABLE [dbo].[Users] ADD  DEFAULT (getdate()) FOR [created_at]
GO
ALTER TABLE [dbo].[BADetails]  WITH CHECK ADD FOREIGN KEY([section_id])
REFERENCES [dbo].[BASection] ([id])
GO
ALTER TABLE [dbo].[BAProgress]  WITH CHECK ADD FOREIGN KEY([detail_id])
REFERENCES [dbo].[BADetails] ([id])
GO
ALTER TABLE [dbo].[BASection]  WITH CHECK ADD FOREIGN KEY([ba_id])
REFERENCES [dbo].[BA] ([id])
GO
ALTER TABLE [dbo].[BreakdownItems]  WITH CHECK ADD FOREIGN KEY([section_id])
REFERENCES [dbo].[BreakdownSections] ([section_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[Breakdowns]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[BreakdownSections]  WITH CHECK ADD FOREIGN KEY([breakdown_id])
REFERENCES [dbo].[Breakdowns] ([breakdown_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[Menus]  WITH CHECK ADD  CONSTRAINT [FK_Menus_ParentID] FOREIGN KEY([parent_id])
REFERENCES [dbo].[Menus] ([menu_id])
GO
ALTER TABLE [dbo].[Menus] CHECK CONSTRAINT [FK_Menus_ParentID]
GO
ALTER TABLE [dbo].[OpnameDetail]  WITH CHECK ADD FOREIGN KEY([section_id])
REFERENCES [dbo].[OpnameSection] ([id])
GO
ALTER TABLE [dbo].[OpnameMandor]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
GO
ALTER TABLE [dbo].[OpnamePayment]  WITH CHECK ADD FOREIGN KEY([opname_mandor_id])
REFERENCES [dbo].[OpnameMandor] ([id])
GO
ALTER TABLE [dbo].[OpnameRekap]  WITH CHECK ADD FOREIGN KEY([opname_mandor_id])
REFERENCES [dbo].[OpnameMandor] ([id])
GO
ALTER TABLE [dbo].[OpnameSection]  WITH CHECK ADD FOREIGN KEY([opname_mandor_id])
REFERENCES [dbo].[OpnameMandor] ([id])
GO
ALTER TABLE [dbo].[ProjectAdditionalExpenses]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
GO
ALTER TABLE [dbo].[ProjectCashFlow]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
GO
ALTER TABLE [dbo].[ProjectCashFlowDetails]  WITH CHECK ADD FOREIGN KEY([cashflow_id])
REFERENCES [dbo].[ProjectCashFlow] ([id])
GO
ALTER TABLE [dbo].[ProjectRekap]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
GO
ALTER TABLE [dbo].[ProjectUser]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
GO
ALTER TABLE [dbo].[ProjectUser]  WITH CHECK ADD FOREIGN KEY([user_id])
REFERENCES [dbo].[Users] ([user_id])
GO
ALTER TABLE [dbo].[RekapPengeluaranDetails]  WITH CHECK ADD FOREIGN KEY([section_id])
REFERENCES [dbo].[RekapPengeluaranSections] ([section_id])
GO
ALTER TABLE [dbo].[RekapPengeluaranSections]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
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
ALTER TABLE [dbo].[Sph]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[SphDetails]  WITH CHECK ADD FOREIGN KEY([section_id])
REFERENCES [dbo].[SphSections] ([section_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[SphSections]  WITH CHECK ADD FOREIGN KEY([sph_id])
REFERENCES [dbo].[Sph] ([sph_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[Spk]  WITH CHECK ADD FOREIGN KEY([project_id])
REFERENCES [dbo].[Projects] ([project_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[SPKDetails]  WITH CHECK ADD FOREIGN KEY([section_id])
REFERENCES [dbo].[SPKSections] ([section_id])
ON DELETE CASCADE
GO
ALTER TABLE [dbo].[SPKSections]  WITH CHECK ADD FOREIGN KEY([spk_id])
REFERENCES [dbo].[Spk] ([spk_id])
ON DELETE CASCADE
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
