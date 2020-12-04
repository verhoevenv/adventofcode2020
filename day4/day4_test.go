package main

import "testing"

var batch = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestParse(t *testing.T) {

	ps := parse(batch)

	if len(ps) != 4 {
		t.Errorf("Parse incorrect, got: %v passports, want %v.",
			len(ps), 4)
	}

	if len(ps[0].fields) != 8 {
		t.Errorf("Parse incorrect, first passport has %v fields, want %v.",
			len(ps[0].fields), 8)
	}

	if ps[2].fields["eyr"] != "2024" {
		t.Errorf("Parse incorrect, third passport field eyr is %v, want %v.",
			ps[2].fields["eyr"], "2024")
	}
}

func TestValidOld(t *testing.T) {

	ps := parse(batch)

	tables := []struct {
		passID int
		valid  bool
	}{
		{0, true},
		{1, false},
		{2, true},
		{3, false},
	}

	for _, table := range tables {
		result := validOld(&ps[table.passID])
		if result != table.valid {
			t.Errorf("Validation of passport %v was incorrect, got: %v, want: %v.",
				table.passID, result, table.valid)
		}
	}
}

var invalidPassportsBatch = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:194cm`

var validPassportsBatch = `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

func TestValidStrict(t *testing.T) {
	invalidPassports := parse(invalidPassportsBatch)
	validPassports := parse(validPassportsBatch)

	for i, pass := range invalidPassports {
		if validStrict(&pass) {
			t.Errorf("Passport %v should not be valid", i)
		}
	}

	for i, pass := range validPassports {
		if !validStrict(&pass) {
			t.Errorf("Passport %v should not be invalid", i)
		}
	}
}
