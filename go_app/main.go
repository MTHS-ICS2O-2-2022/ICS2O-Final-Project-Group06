// Created by: Jaden Plugowsky, Dominik Armatys
// Created on: June 2023
//
// This program is a short text adventure with a nonlinear path

package main

import "fmt"

var currentAnswer float64

func main() {
	// This function is a short text adventure with a nonlinear path
	fmt.Println("----------------------------------------------------------")
	fmt.Println("		Welcome to Text Adventure!")
	fmt.Println("\nIn this game, you will progress down a short adventure with a nonlinear path.")
	fmt.Println("\nAnswer with the numbers '1' or '2' for the first and second choices, respectively.")
	fmt.Println("\nType number '1' or number '2' to begin!")
	currentAnswer = 0
	fmt.Scanln(&currentAnswer)
	if currentAnswer == 1 || currentAnswer == 2 { // Start
		fmt.Println("\nGood job, you now know how to play!")
		fmt.Println("And now, we begin.")
		fmt.Println("\n----------------------------------------------------------")
		fmt.Println("\nYou find yourself in a forest, and you're lost.")
		fmt.Println("\nDo you..\nTalk to dom (1)\nor\nGo straight ahead (2)")
		currentAnswer = 0
		fmt.Scanln(&currentAnswer)
		if currentAnswer == 1 {
			fmt.Println("\n----------------------------------------------------------")
			fmt.Println("\nYou talk to dom.")
			fmt.Println("\nHe tells you to go to the nearby lake and go for a swim.\nHe doesn't elaborate.")
			fmt.Println("\nDo you..\nInterrogate dom (1)\nor\nGo to the lake and swim (2)")
			currentAnswer = 0
			fmt.Scanln(&currentAnswer)
			if currentAnswer == 1 {
				fmt.Println("\n----------------------------------------------------------")
				fmt.Println("\nYou interrogate dom.")
				fmt.Println("\nHe says the lake is super cold and would have killed you if you went for a swim.")
				fmt.Println("\nThat's pretty mean of him, do you..\nScream at dom for suggesting you die (1)\nor\nShrug and say 'that's okay, man' (2)")
				currentAnswer = 0
				fmt.Scanln(&currentAnswer)
				if currentAnswer == 1 {
					fmt.Println("\n----------------------------------------------------------")
					fmt.Println("\nYou scream various not nice words at dom.")
					fmt.Println("\nSuddenly, he looks really sad. \nLike, super sad. \nHe might be the saddest he has ever been.")
					fmt.Println("\nYou monster. Do you..\nSay sorry (1)\nor\nTurn away from him in a rude way (2)")
					currentAnswer = 0
					fmt.Scanln(&currentAnswer)
					if currentAnswer == 1 {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\nYou say sorry to dom for yelling at him.")
						fmt.Println("\nHe doesn't accept your apology (rightfully) and simply walks away, leaving you to fend for yourself alone in the forest.")
						fmt.Println("\nYOU ARE LOST")
						fmt.Println("\nPut in the command to try again!")
					} else if currentAnswer == 2 {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\nYou turn away from dom in a rude and mean way.\ndom is still crying, which makes you a mean person.")
						fmt.Println("\nYou simply walk away in the opposite direction, which is nothing but dense bushes for days.\nSomehow, you are more lost than when you started.")
						fmt.Println("\nYOU ARE LOST")
						fmt.Println("\nPut in the command to try again!")
					} else {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\nThat input is invalid, please play again.")
					}
				} else if currentAnswer == 2 {
					fmt.Println("\n----------------------------------------------------------")
					fmt.Println("\nYou tell dom it's alright that he suggested you swim in the cold lake which would have absolutely killed you.")
					fmt.Println("\nHe holds out both his closed hands, saying, 'cool thanks, here's a lil something for forgiving me.'")
					fmt.Println("\nWhich hand do you choose?\ndom's left hand (1)\nor\ndom's right hand (2)")
					currentAnswer = 0
					fmt.Scanln(&currentAnswer)
					if currentAnswer == 1 {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\ndom opens both his left and right hands.")
						fmt.Println("\nHis right hand is empty, but the left one has a cartoon bomb with a very short (and burning) fuse.")
						fmt.Println("You don't see why dom did this, but you think it's because he has been super bored during his time in this forest.\nAnyways, the bomb exploded, but dom's okay, so that's something.\n\nYou, however, are not.")
						fmt.Println("\nYOU DIED")
						fmt.Println("\nAnd, so close to the ending, too. Put in the command to try again!")
					} else if currentAnswer == 2 {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\nYou choose dom's right hand.")
						fmt.Println("\nHe gives you a map, showing that his other hand was empty the whole time.")
						fmt.Println("\nYou look at the map, noticing that it shows you a specific and direct path from your current location to out of the forest.")
						fmt.Println("\nYOU ESCAPED!")
						fmt.Println("\nBravo! You found one of the two correct paths, and you escaped the forest!\nNext time, go straight ahead instead of talking to dom, and you'll be on track to the other good ending!.")
					} else {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\nThat input is invalid, please play again.")
					}
				} else {
					fmt.Println("\n----------------------------------------------------------")
					fmt.Println("\nThat input is invalid, please play again.")
				}
			} else if currentAnswer == 2 {
				fmt.Println("\n----------------------------------------------------------")
				fmt.Println("\nYou go to the lake and take a swim.")
				fmt.Println("\nYou freeze in the extremely and surprisingly cold water.\ndom had forsaken you.")
				fmt.Println("\nYOU DIED")
				fmt.Println("\nPut in the command to try again!")
			} else {
				fmt.Println("\n----------------------------------------------------------")
				fmt.Println("\nThat input is invalid, please play again.")
			}
		} else if currentAnswer == 2 {
			fmt.Println("\n----------------------------------------------------------")
			fmt.Println("\nYou walk ahead and find an old, creepy house.")
			fmt.Println("\nYou are suddenly hungry, and you wonder if you could find food in the house.")
			fmt.Println("\nDo you..\nSearch for food in the house (1)\nor\nAvoid the creepy house and keep walking (2)")
			currentAnswer = 0
			fmt.Scanln(&currentAnswer)
			if currentAnswer == 1 {
				fmt.Println("\n----------------------------------------------------------")
				fmt.Println("\nYou walk into the creepy house in search of food.")
				fmt.Println("\nA witch suddenly appears and kills you.")
				fmt.Println("\nYOU DIED")
				fmt.Println("\nPut in the command to try again!")
			} else if currentAnswer == 2 {
				fmt.Println("\n----------------------------------------------------------")
				fmt.Println("\nYou keep walking.")
				fmt.Println("\nEventually, you find a pile of beautiful, glowing, and probaby very deadly mushrooms.")
				fmt.Println("\nDo you..\nEat a mushroom (1)\nor\nNo, I'd rather live, to be honest (2)")
				currentAnswer = 0
				fmt.Scanln(&currentAnswer)
				if currentAnswer == 1 {
					fmt.Println("\n----------------------------------------------------------")
					fmt.Println("\nYou eat one of the glowing, probably radioactive and likely super-dangerous mushrooms.")
					fmt.Println("\nSurprisingly, it doesn't kill you, and you feel a lot less hungry.\nInstead, you feel as though you could move much faster than before.")
					fmt.Println("\nDo you..\nRun as quickly as possible out of the forest (1)\nor\nRun towards the lake (2)")
					currentAnswer = 0
					fmt.Scanln(&currentAnswer)
					if currentAnswer == 1 {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\nYou begin running as quickly as possible in a random direction in an attempt to escape the forest.")
						fmt.Println("\nUnfortunately, you don't know where to go to escape.")
						fmt.Println("Even worse: you start heating up due to friction or physics or something, and you explode in a large fireball.")
						fmt.Println("You also start a forest fire, which wasn't very cool of you.")
						fmt.Println("\nYOU DIED")
						fmt.Println("\nAnd, so close to the ending, too. Put in the command to try again!")
					} else if currentAnswer == 2 {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\nYou begin running as quickly as possible in the direction of the lake.")
						fmt.Println("\nYou notice yourself heating up, and you would have died if you hadn't started running on top of the lake, which was somehow cold enough to stop you from overheating.")
						fmt.Println("You run to the other side of the lake, finally free from the forest.")
						fmt.Println("\nYOU ESCAPED!")
						fmt.Println("\nGood job! You have found one of the two correct paths, and you escaped the forest!\nDid you know to run on the super-cold lake from a previous death?")
						fmt.Println("Next time, you should talk to dom instead of going straight ahead, and then you'll be on track to get the other good ending!")
					} else {
						fmt.Println("\n----------------------------------------------------------")
						fmt.Println("\nThat input is invalid, please try again!")
					}
				} else if currentAnswer == 2 {
					fmt.Println("\n----------------------------------------------------------")
					fmt.Println("\nYou choose not to eat the mushroom that, under pretty much every circumstance, would make you die a very slow and agonizing death due to poison or some other horrible thing.")
					fmt.Println("\nHowever, you suddenly notice how hungry you are, and, having found no food, you drop to the ground, weak and starving.")
					fmt.Println("\nYOU DIED")
					fmt.Println("\nPut in the command to try again!")
				} else {
					fmt.Println("\n----------------------------------------------------------")
					fmt.Println("\nThat input is invalid, please play again.")
				}
			} else {
				fmt.Println("\n----------------------------------------------------------")
				fmt.Println("\nThat input is invalid, please play again.")
			}
		} else {
			fmt.Println("\n----------------------------------------------------------")
			fmt.Println("\nThat input is invalid, please play again.")
		}
	} else {
		fmt.Println("\n----------------------------------------------------------")
		fmt.Println("\nThat input is invalid, please play again.")
	}
	fmt.Println("\n----------------------------------------------------------")

	fmt.Println("\nDone.")
}
