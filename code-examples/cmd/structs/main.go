package main

import "fmt"

type details struct {
	name      string
	dob       string
	interests string
	jobstuff
}

type jobstuff struct {
	job     string
	company string
	number  int
}

func (j jobstuff) jobDetails() string {
	return fmt.Sprintf("%s, %s, %d\n", j.job, j.company, j.number)
}

func (j *jobstuff) changeJobDetails(choice string) {
	switch choice {
	case "job":
		var setJob string
		fmt.Print("Enter new job title: ")
		fmt.Scanln(&setJob)
		j.job = setJob

	case "company":
		var setCompany string
		fmt.Print("Enter new company name: ")
		fmt.Scanln(&setCompany)
		j.job = setCompany

	case "number":
		var setNumber string
		fmt.Print("Enter new number: ")
		fmt.Scanln(&setNumber)
		j.job = setNumber

	default:
		fmt.Println("Invalid choice. Valid options are: job, company, number.")
	}
}

func main() {
	var name, dob, interests, job, company string
	var number int

	fmt.Print("Enter you name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter you dob (DD/MM/YYYY): ")
	fmt.Scanln(&dob)

	fmt.Print("Enter an interest of yours: ")
	fmt.Scanln(&interests)

	fmt.Print("Enter you job title: ")
	fmt.Scanln(&job)

	fmt.Print("Enter you company name: ")
	fmt.Scanln(&company)

	fmt.Print("Enter you phone number: ")
	fmt.Scanln(&number)

	p := details{
		name:      name,
		dob:       dob,
		interests: interests,
		jobstuff: jobstuff{
			job:     job,
			company: company,
			number:  number,
		},
	}

	fmt.Printf("%s, %s, %s, %s\n", p.name, p.dob, p.interests, p.jobDetails())

	var choice string
	fmt.Print("Would you like to update any of the job details? Enter job/company/number/no: ")
	fmt.Scanln(&choice)

	if choice != "none" {
		p.changeJobDetails((choice))
	}
	fmt.Printf("Final Details: %s, %s, %s, %s\n", p.name, p.dob, p.interests, p.jobDetails())
}

/* Create a command line application (with unit testing) that allows a user to enter their personal details (make them up!):
Name
Date of Birth
Interests
Job Title
Contact Details

You can call the fields what you like, and have as many fields as you think are necessary, but you must show at least the following:
Use of structs DONE
Use of composition (embedding) DONE
Use of receivers

Your application should allow a user to:
Enter their details DONE
Change any of their details
Print out their details */
