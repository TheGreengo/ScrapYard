#!/usr/bin/perl

my $oi;
my $vei = 4;

$oi //= 3;
$vei //= 23;

print "$oi\n";
print "$vei\n";

print `pwd`;
print "\n";
print `pwd`;

my @things = (1, 2, 3, 5);
my $thang = shift @things;
print "$thang\n";
print @things;
