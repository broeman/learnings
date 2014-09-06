using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SticksGame
{
    class HumanPlayer : Player
    {
        public HumanPlayer()
        {
            Console.Write("Hvad er dit navn? ");
            Name = Console.ReadLine();         
            PlayerType = "Human";
        }

        public override int GetMove(int sticksInPlay)
        {
            int numVal = 0;
            bool repeat = true;
        Console.Write("Hvor mange tager du (1,2 eller 3)? ");
            string userinput;
            while (repeat)
            {
                try
                {
                    repeat = true;
                    userinput = Console.ReadLine();
                    numVal = Convert.ToInt32(userinput);
                    if (numVal < 1 || numVal > 3)
                    {
                        Console.WriteLine("\nDu kan kun tage 1 til 3 pinde.\n");
                    }
                    else
                    {
                        repeat = false;
                    }
                }
                catch (FormatException)
                {
                    Console.WriteLine("Det skal være et tal");
                }
                catch (OverflowException)
                {
                    Console.WriteLine("It cannot fit in an Integer 32-bit");
                }
            }
            return numVal;
        }
    }
}
