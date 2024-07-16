<?php

require_once __DIR__ . '/vendor/autoload.php';

use GameInspire\CodeStyle\PlatformConfig;
use PhpCsFixer\Finder;

$finder = Finder::create()
    ->in(
        [
            __DIR__ . '/src',
            __DIR__ . '/tests',
        ]
    );

return (new PhpCsFixer\Config())
    ->setFinder($finder)
    ->registerCustomFixers(new GameInspire\CodeStyle\Fixers())
    ->setRules(PlatformConfig::getRules());
