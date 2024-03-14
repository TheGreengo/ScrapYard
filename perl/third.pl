#!/usr/bin/perl
use strict;
use warnings;

my %thing = (
    word => 1,
    werd => 2
);

my %thang = qw(
    o word
    e werd
);

my $one;
my $two;

chomp($one = <STDIN>);
chomp($two = <STDIN>);

$thang{$one} = $two;

delete $thang{'o'};

for (keys %thang) {
    print("The key $_ corresponds to the entry: $thang{$_}\n");
}