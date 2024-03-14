#!/usr/bin/perl
use strict;
use warnings;

my %thing = (
    word => 1,
    werd => 2,
);

my $res;
chomp($res = <STDIN>);

print($thing{$res}, "\n");