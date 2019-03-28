<?php

require_once '../vendor/autoload.php';

use PhpAmqpLib\Message\AMQPMessage;
use PhpAmqpLib\Connection\AMQPStreamConnection;

// CloudAMQP configuration
$host = 'crane.rmq.cloudamqp.com';
$port = 5672;
$user = 'xxxxx';
$password = 'xxxxx';
$vhost = 'xxxxx';

// Initiated connection to CloudAMQP
$conn = new AMQPStreamConnection($host, $port, $user, $password, $vhost);
$channel = $conn->channel();

// Queue name, this name must be same with consumer.
$queueName = 'profile';
$channel->queue_declare($queueName, false, false, false, false);

// Initiate message to be send to consumer
$txt = 'My name is Mohd Norlihazmey Ghazali';
$msg = new AMQPMessage($txt);
$channel->basic_publish($msg, '', $queueName);

echo 'Message has been sent and place inside queue.';

// Close channel and connection
$channel->close();
$conn->close();
