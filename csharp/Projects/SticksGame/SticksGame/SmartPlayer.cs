using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SticksGame
{
    class SmartPlayer : RandomPlayer
    {
        public SmartPlayer()
        {
            Name = "Smart";
        }

        public override int GetMove(int sticksInPlay)
        {
            Random result = new Random();
            switch (sticksInPlay) {
                // logic: 4*n + 1 => 1: 5, 2: 9, 3: 13
                // return (sticksInPlay - 1) % 4;
                case 2:
                    return 1;
                case 3:
                    return 2;
                case 4:
                    return 3;
                case 6:
                    return 1;
                case 7:
                    return 2;
                case 8:
                    return 3;
                case 10:
                    return 1;
                case 11:
                    return 2;
                case 12:
                    return 3;
                default:
                    return base.GetMove(sticksInPlay);
            }
        }
    }
}
