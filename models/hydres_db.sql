CREATE TABLE `series`
(
  `id` int PRIMARY KEY,
  `name` string,
  `status` orders_status,
  `startet_at` datetime,
  `ended_at` datetime
);

CREATE TABLE `seed`
(
  `id` int PRIMARY KEY,
  `research_tag` string,
  `series_id` int UNIQUE NOT NULL,
  `name` string,
  `created_at` datetime,
  `type` string
);

CREATE TABLE `seed_journal`
(
  `id` int PRIMARY KEY,
  `seed_id` int UNIQUE NOT NULL,
  `research_tag` string,
  `name` string,
  `created_at` datetime,
  `text` string,
  `photo` string
);

CREATE TABLE `plant`
(
  `id` int PRIMARY KEY,
  `series_id` int,
  `research_tag` string,
  `name` string,
  `replant_at` datetime,
  `type` string
);

CREATE TABLE `plant_journal`
(
  `id` int PRIMARY KEY,
  `plant_id` int,
  `research_tag` string,
  `name` string,
  `created_at` datetime,
  `text` string,
  `photo` string
);

ALTER TABLE `seed` ADD FOREIGN KEY (`series_id`) REFERENCES `series` (`id`);

ALTER TABLE `seed_journal` ADD FOREIGN KEY (`research_tag`) REFERENCES `seed` (`research_tag`);

ALTER TABLE `seed_journal` ADD FOREIGN KEY (`seed_id`) REFERENCES `seed` (`id`);

ALTER TABLE `plant` ADD FOREIGN KEY (`series_id`) REFERENCES `series` (`id`);

ALTER TABLE `plant` ADD FOREIGN KEY (`research_tag`) REFERENCES `seed` (`research_tag`);

ALTER TABLE `plant_journal` ADD FOREIGN KEY (`plant_id`) REFERENCES `plant` (`id`);

ALTER TABLE `plant_journal` ADD FOREIGN KEY (`research_tag`) REFERENCES `seed` (`research_tag`);

