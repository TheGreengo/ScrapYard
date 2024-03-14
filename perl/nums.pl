#!/usr/bin/perl
use strict;
use warnings;

my $x = 123_456_789;
my $inp;
my @arr1 = qw(thing thang theng);
my @arr2 = (1, 4, 16);

chomp($inp = <STDIN>);

print($x, "\n");
print($inp, "\n");
print($arr1[0], "\n");
print($arr2[2], "\n");