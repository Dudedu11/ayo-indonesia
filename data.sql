-- team data
INSERT INTO public.teams
(id, created_at, updated_at, deleted_at, "name", logo_url, founded_year, base_address, base_city)
VALUES(1, '2025-08-12 22:36:42.805', '2025-08-12 22:36:42.805', NULL, 'Persib Bandung', 'Persib.png', 1933, 'Jl. Sulanjana No.17, Tamansari', 'Bandung');
INSERT INTO public.teams
(id, created_at, updated_at, deleted_at, "name", logo_url, founded_year, base_address, base_city)
VALUES(2, '2025-08-12 22:37:14.924', '2025-08-12 22:37:14.924', NULL, 'Persija Jakarta', 'Persija.png', 1928, 'Jl. Raya Matraman No.19, Jakarta Timur', 'Jakarta');
INSERT INTO public.teams
(id, created_at, updated_at, deleted_at, "name", logo_url, founded_year, base_address, base_city)
VALUES(3, '2025-08-12 22:37:32.490', '2025-08-12 22:37:32.490', NULL, 'Persebaya Surabaya', 'Persebaya.png', 1927, 'Jl. Karanggayam No.1, Tambaksari', 'Surabaya');
INSERT INTO public.teams
(id, created_at, updated_at, deleted_at, "name", logo_url, founded_year, base_address, base_city)
VALUES(4, '2025-08-12 22:37:45.534', '2025-08-12 22:37:45.534', NULL, 'Arema FC', 'Arema.png', 1987, 'Jl. Mayjen Pandjaitan No.42, Blimbing', 'Malang');

-- player data
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(1, '2025-08-12 22:44:38.072', '2025-08-12 22:44:38.072', NULL, 'Adam Przybek', 185, 78, 'penjaga gawang', 1, 1);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(2, '2025-08-12 22:44:46.588', '2025-08-12 22:44:46.588', NULL, 'Patricio Matricardi', 188, 83, 'bertahan', 48, 1);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(3, '2025-08-12 22:44:54.018', '2025-08-12 22:44:54.018', NULL, 'Marc Klok', 177, 72, 'gelandang', 23, 1);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(4, '2025-08-12 22:45:02.535', '2025-08-12 22:45:02.535', NULL, 'Dimas Drajad', 179, 75, 'penyerang', 9, 1);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(5, '2025-08-12 22:45:12.774', '2025-08-12 22:45:12.774', NULL, 'Andritany Ardhiyasa', 178, 76, 'penjaga gawang', 26, 2);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(6, '2025-08-12 22:45:19.223', '2025-08-12 22:45:19.223', NULL, 'Rizky Ridho Ramadhani', 183, 78, 'bertahan', 5, 2);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(7, '2025-08-12 22:45:26.705', '2025-08-12 22:45:26.705', NULL, 'Gustavo França', 175, 71, 'gelandang', 10, 2);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(8, '2025-08-12 22:45:33.278', '2025-08-12 22:45:33.278', NULL, 'Gustavo Almeida dos Santos', 180, 76, 'penyerang', 70, 2);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(9, '2025-08-12 22:45:41.410', '2025-08-12 22:45:41.410', NULL, 'Muhammad Ferrari', 180, 74, 'penjaga gawang', 30, 3);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(10, '2025-08-12 22:45:56.203', '2025-08-12 22:45:56.203', NULL, 'Slavko Damjanović', 190, 85, 'bertahan', 4, 3);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(11, '2025-08-12 22:46:02.489', '2025-08-12 22:46:02.489', NULL, 'Firza Andika', 178, 70, 'gelandang', 11, 3);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(12, '2025-08-12 22:46:08.902', '2025-08-12 22:46:08.902', NULL, 'Paulo Henrique Leo Lelis', 182, 76, 'penyerang', 9, 3);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(13, '2025-08-12 22:46:16.048', '2025-08-12 22:46:16.048', NULL, 'Adi Satryo', 185, 79, 'penjaga gawang', 30, 4);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(14, '2025-08-12 22:46:23.735', '2025-08-12 22:46:23.735', NULL, 'Odivan Koerich', 188, 84, 'bertahan', 4, 4);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(15, '2025-08-12 22:46:30.813', '2025-08-12 22:46:30.813', NULL, 'Valdeci', 175, 70, 'gelandang', 10, 4);
INSERT INTO public.players
(id, created_at, updated_at, deleted_at, "name", height, weight, "position", "number", team_id)
VALUES(16, '2025-08-12 22:46:38.718', '2025-08-12 22:46:38.718', NULL, 'Dimas Aryaguna', 180, 74, 'penyerang', 70, 4);

-- match data 
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(1, '2025-08-12 22:52:57.015', '2025-08-12 22:52:57.015', NULL, '2025-08-10 22:30:00.000', '15:30', 1, 2);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(2, '2025-08-12 22:53:05.305', '2025-08-12 22:53:05.305', NULL, '2025-08-13 23:00:00.000', '16:00', 1, 3);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(3, '2025-08-12 22:53:13.063', '2025-08-12 22:53:13.063', NULL, '2025-08-17 01:00:00.000', '18:00', 1, 4);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(4, '2025-08-12 22:53:21.084', '2025-08-12 22:53:21.084', NULL, '2025-08-19 22:30:00.000', '15:30', 2, 3);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(5, '2025-08-12 22:53:26.122', '2025-08-12 22:53:26.122', NULL, '2025-08-23 00:00:00.000', '17:00', 2, 4);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(6, '2025-08-12 22:53:38.244', '2025-08-12 22:53:38.244', NULL, '2025-08-26 03:00:00.000', '20:00', 3, 4);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(7, '2025-08-13 20:29:42.628', '2025-08-13 20:29:42.628', NULL, '2025-09-10 07:00:00.000', '19:00', 2, 1);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(8, '2025-08-13 20:29:53.219', '2025-08-13 20:29:53.219', NULL, '2025-09-17 07:00:00.000', '19:00', 3, 1);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(9, '2025-08-13 20:30:01.066', '2025-08-13 20:30:01.066', NULL, '2025-09-24 07:00:00.000', '19:00', 4, 1);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(10, '2025-08-13 20:30:19.106', '2025-08-13 20:30:19.106', NULL, '2025-10-08 07:00:00.000', '19:00', 3, 2);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(11, '2025-08-13 20:30:26.953', '2025-08-13 20:30:26.953', NULL, '2025-10-15 07:00:00.000', '19:00', 4, 2);
INSERT INTO public.matches
(id, created_at, updated_at, deleted_at, "date", "time", home_team_id, away_team_id)
VALUES(12, '2025-08-13 20:30:49.334', '2025-08-13 20:30:49.334', NULL, '2025-11-05 07:00:00.000', '19:00', 4, 3);

-- match result data 
INSERT INTO public.match_results
(id, created_at, updated_at, deleted_at, match_id, home_score, away_score, winner_team_id)
VALUES(1, '2025-08-13 20:44:57.975', '2025-08-13 20:45:27.933', NULL, 1, 5, 0, 1);
INSERT INTO public.match_results
(id, created_at, updated_at, deleted_at, match_id, home_score, away_score, winner_team_id)
VALUES(2, '2025-08-13 21:18:32.817', '2025-08-13 21:18:32.817', NULL, 2, 0, 0, NULL);

-- goal data 
INSERT INTO public.goal_details
(id, created_at, updated_at, deleted_at, match_id, player_id, team_id, "minute")
VALUES(1, '2025-08-13 20:44:57.956', '2025-08-13 20:44:57.956', NULL, 1, 4, 1, 12);
INSERT INTO public.goal_details
(id, created_at, updated_at, deleted_at, match_id, player_id, team_id, "minute")
VALUES(2, '2025-08-13 20:45:07.134', '2025-08-13 20:45:07.134', NULL, 1, 4, 1, 28);
INSERT INTO public.goal_details
(id, created_at, updated_at, deleted_at, match_id, player_id, team_id, "minute")
VALUES(3, '2025-08-13 20:45:13.808', '2025-08-13 20:45:13.808', NULL, 1, 3, 1, 41);
INSERT INTO public.goal_details
(id, created_at, updated_at, deleted_at, match_id, player_id, team_id, "minute")
VALUES(4, '2025-08-13 20:45:19.882', '2025-08-13 20:45:19.882', NULL, 1, 2, 1, 55);
INSERT INTO public.goal_details
(id, created_at, updated_at, deleted_at, match_id, player_id, team_id, "minute")
VALUES(5, '2025-08-13 20:45:27.931', '2025-08-13 20:45:27.931', NULL, 1, 4, 1, 77);
