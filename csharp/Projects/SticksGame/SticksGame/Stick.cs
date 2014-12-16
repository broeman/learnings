using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace SticksGame
{
    class Stick
    {
        public int Count { get; set; }
        public Stick(int numSticks) {
            Count = numSticks;
        }
    }
}
