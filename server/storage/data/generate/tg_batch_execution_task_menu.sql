-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和在SQL文件：D:\main\grata\server\storage\data\generate\tg_batch_execution_task_menu.sql


SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;

--
-- 数据库： `grata`
--

-- --------------------------------------------------------

--
-- 插入表中的数据 `hg_admin_menu`
--


SET @now := now();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '62076', '批量操作任务', 'tgBatchExecutionTaskIndex', 'tgBatchExecutionTask/index', 'MenuOutlined', '2', '', '/tgBatchExecutionTask/list', '', '/tg/tgBatchExecutionTask/index', '1', '', '0', '0', '', '0', '0', '0', '2', '', '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '62076', '批量操作任务详情', 'tgBatchExecutionTaskView', 'tgBatchExecutionTask/view/:id?', '', '2', '', '/tgBatchExecutionTask/view', '', '/tg/tgBatchExecutionTask/view', '0', 'tgBatchExecutionTaskIndex', '0', '0', '', '0', '1', '0', '2', '', '20', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增批量操作任务', 'tgBatchExecutionTaskEdit', '', '', '3', '', '/tgBatchExecutionTask/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '3', '', '10', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除批量操作任务', 'tgBatchExecutionTaskDelete', '', '', '3', '', '/tgBatchExecutionTask/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '3', '', '10', '', '1', @now, @now);


-- 更新状态
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '修改批量操作任务状态', 'tgBatchExecutionTaskStatus', '', '', '3', '', '/tgBatchExecutionTask/status', '', '', '1', '', '0', '0', '', '0', '0', '0', '3', '', '10', '', '1', @now, @now);



-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '导出批量操作任务', 'tgBatchExecutionTaskExport', '', '', '3', '', '/tgBatchExecutionTask/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '3', '', '10', '', '1', @now, @now);


COMMIT;