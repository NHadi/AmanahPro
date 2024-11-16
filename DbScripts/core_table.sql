USE [AmanahDb]
GO
/****** Object:  Table [dbo].[BA]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[BADetails]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[BAProgress]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[BASection]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[BreakdownItems]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[Breakdowns]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[BreakdownSections]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[Invoices]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[OpnameDetail]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[OpnameMandor]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[OpnamePayment]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[OpnameRekap]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[OpnameSection]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[ProjectAdditionalExpenses]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[ProjectCashFlow]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[ProjectCashFlowDetails]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[ProjectRekap]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[Projects]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[ProjectUser]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[RekapPengeluaranDetails]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[RekapPengeluaranSections]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[Sph]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[SphDetails]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[SphSections]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[Spk]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[SPKDetails]    Script Date: 14/11/2024 23:18:25 ******/
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
/****** Object:  Table [dbo].[SPKSections]    Script Date: 14/11/2024 23:18:25 ******/
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
