ALTER TABLE `metaspace`.`contract_trace_log`
    ADD COLUMN `chain` varchar(192) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL AFTER `event_uuid`;