#!/usr/bin/perl
use strict;
use warnings;
use v5.10;

my $true = 1;
print("Truth\n") if($true);
print("Falsehood\n") if($true - 1);

if (not $true - 1) {
    print("Truth\n");
} elsif($true - 2) {
    print("Whathood\n");
} else {
    print("Falsehood\n");
}

unless ("aklf") {
    print("Unless is working\n");
} elsif ("asd") {
    print("elsif\n");
} else {
    print("else\n");
}

my $next;
chomp($next = <STDIN>);

given ($next) {
    when($next < 0) {
        print("negative\n");
    }
    when($next > 0) {
        print("positive\n");
    }
    default {
        print("zero\n");
    }
}