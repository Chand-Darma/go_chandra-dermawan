-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Mar 20, 2023 at 08:51 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop`
--

-- --------------------------------------------------------

--
-- Table structure for table `addresses`
--

CREATE TABLE `addresses` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `address_line_1` varchar(255) DEFAULT NULL,
  `address_line_2` varchar(255) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  `state` varchar(100) DEFAULT NULL,
  `country` varchar(100) DEFAULT NULL,
  `postcode` varchar(10) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'operator1', '2023-03-16 18:38:26', '2023-03-16 18:38:26'),
(2, 'operator2', '2023-03-16 18:38:26', '2023-03-16 18:38:26'),
(3, 'operator3', '2023-03-16 18:38:26', '2023-03-16 18:38:26'),
(4, 'operator4', '2023-03-16 18:38:26', '2023-03-16 18:38:26'),
(5, 'operator5', '2023-03-16 18:38:26', '2023-03-16 18:38:26');

-- --------------------------------------------------------

--
-- Table structure for table `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Credit Card', 1, '2023-03-17 07:33:54', '2023-03-17 07:33:54'),
(2, 'Paypal', 2, '2023-03-17 07:34:17', '2023-03-17 07:34:17'),
(3, 'Bank Transfer', 3, '2023-03-17 07:34:43', '2023-03-17 07:34:43');

-- --------------------------------------------------------

--
-- Table structure for table `payment_method_descriptions`
--

CREATE TABLE `payment_method_descriptions` (
  `payment_method_id` int(11) NOT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `product_type_id` int(11) DEFAULT NULL,
  `operator_id` int(11) NOT NULL,
  `code` varchar(50) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `product_description` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`, `product_description`) VALUES
(3, 2, 1, '299', 'Clothing 2', 22, '2023-03-16 18:58:44', '2023-03-17 07:30:38', 'Ready'),
(4, 2, 1, '201', 'Clothing 3', 10, '2023-03-16 18:59:30', '2023-03-17 07:30:48', 'Not Ready'),
(5, 2, 1, '421', 'Clothing 4', 12, '2023-03-16 19:00:39', '2023-03-17 07:30:54', 'Ready'),
(6, 3, 4, '876', 'Home and Kitchen 1', 44, '2023-03-16 19:03:18', '2023-03-17 07:31:01', 'Ready'),
(7, 3, 4, '899', 'Home and Kitchen 2', 14, '2023-03-16 19:03:41', '2023-03-17 07:31:16', 'Not Ready'),
(8, 3, 4, '276', 'Home and Kitchen 3', 64, '2023-03-16 19:04:01', '2023-03-17 07:31:23', 'Ready');

-- --------------------------------------------------------

--
-- Table structure for table `product_descriptions`
--

CREATE TABLE `product_descriptions` (
  `id` int(11) NOT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `product_types`
--

CREATE TABLE `product_types` (
  `id` int(11) NOT NULL,
  `name` varchar(225) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product_types`
--

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Electronics', '2023-03-16 18:40:32', '2023-03-16 18:40:32'),
(2, 'Clothing', '2023-03-16 18:40:32', '2023-03-16 18:40:32'),
(3, 'Home and Kitchen', '2023-03-16 18:40:32', '2023-03-16 18:40:32');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `payment_method_id` int(11) DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `total_qty` int(11) DEFAULT NULL,
  `total_price` decimal(25,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'pending', 3, '1500.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(2, 1, 2, 'paid', 2, '1000.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(3, 1, 3, 'cancelled', 1, '500.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(4, 2, 1, 'pending', 2, '1000.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(5, 2, 2, 'paid', 4, '2000.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(6, 2, 3, 'cancelled', 1, '500.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(7, 3, 1, 'pending', 5, '2500.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(8, 3, 2, 'paid', 3, '1500.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(9, 3, 3, 'cancelled', 2, '1000.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(10, 4, 1, 'pending', 1, '500.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(11, 4, 2, 'paid', 2, '1000.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(12, 4, 3, 'cancelled', 1, '500.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(13, 5, 1, 'pending', 4, '2000.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(14, 5, 2, 'paid', 5, '2500.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57'),
(15, 5, 3, 'cancelled', 2, '1000.00', '2023-03-17 07:42:57', '2023-03-17 07:42:57');

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(10) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `price` decimal(25,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `status` smallint(6) DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 1, '1990-01-01', 'M', '2023-03-17 07:39:52', '2023-03-17 07:39:52'),
(2, 2, '1985-03-15', 'F', '2023-03-17 07:39:52', '2023-03-17 07:39:52'),
(3, 3, '1995-12-31', 'M', '2023-03-17 07:39:52', '2023-03-17 07:39:52'),
(4, 4, '2000-06-12', 'M', '2023-03-17 07:39:52', '2023-03-17 07:39:52'),
(5, 5, '1988-11-23', 'M', '2023-03-17 07:39:52', '2023-03-17 07:39:52');

-- --------------------------------------------------------

--
-- Table structure for table `user_payment_method_detail_history`
--

CREATE TABLE `user_payment_method_detail_history` (
  `id` int(11) NOT NULL,
  `user_payment_method_detail_id` int(11) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `addresses`
--
ALTER TABLE `addresses`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_method_descriptions`
--
ALTER TABLE `payment_method_descriptions`
  ADD PRIMARY KEY (`payment_method_id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_type_id` (`product_type_id`);

--
-- Indexes for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `payment_method_id` (`payment_method_id`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`transaction_id`,`product_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_payment_method_detail_history`
--
ALTER TABLE `user_payment_method_detail_history`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_payment_method_detail_id` (`user_payment_method_detail_id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `addresses`
--
ALTER TABLE `addresses`
  ADD CONSTRAINT `addresses_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `payment_method_descriptions`
--
ALTER TABLE `payment_method_descriptions`
  ADD CONSTRAINT `payment_method_descriptions_ibfk_1` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`);

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`product_type_id`) REFERENCES `product_types` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`);

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

--
-- Constraints for table `user_payment_method_detail_history`
--
ALTER TABLE `user_payment_method_detail_history`
  ADD CONSTRAINT `user_payment_method_detail_history_ibfk_1` FOREIGN KEY (`user_payment_method_detail_id`) REFERENCES `user_payment_method_detail` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
