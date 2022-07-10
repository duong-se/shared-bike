-- +goose Up
-- SQL in this section is executed when the migration is applied.

INSERT INTO `user` (`id`, `username`, `password`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'test1', '$2a$10$FL6I3Q8Y1OQMjzO7M1HLa.DAQN/IwmeCxgpQr/u8Yq3tissOwVOam', 'Test User', '2022-07-06 18:51:44', '2022-07-06 18:51:44', NULL),
(2, 'test2', '$2a$10$45MrKZFzZo4Ks8utOaHCKuAy1E02Bft7ED92PjsMxcIjg1AyNWhXC', 'Test User 2', '2022-07-06 19:23:43', '2022-07-06 19:23:43', NULL);

INSERT INTO `bike` (`id`, `name`, `lat`, `long`, `status`, `user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, "Henry", 50.123456, 123456, 'available', NULL, '2022-07-06 16:55:25', '2022-07-07 16:20:09', NULL),
(2, "Harry", 50.119504, 8.638137, 'available', NULL, '2022-07-06 18:49:47', '2022-07-07 14:11:37', NULL),
(3, "Copper", 50.119229, 8.640020, 'available', NULL, '2022-07-06 18:50:07', '2022-07-06 18:50:07', NULL),
(4, "Titan", 50.120452, 8.650507, 'available', NULL, '2022-07-06 18:50:07', '2022-07-06 18:50:07', NULL);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
TRUNCATE TABLE `bike`;
TRUNCATE TABLE `user`;
