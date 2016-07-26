# generates songdata.go for use by Gonant

num_insts = 8

def process_instruments(lines):
    idx = 0

    for instr in range(num_insts):
        print("\t\tInstrument" + lines[idx].lstrip())
        idx += 1
        
        while lines[idx].strip() != "// Patterns" :
            print(lines[idx])
            idx += 1

        print(lines[idx])
        idx += 1
        print("\t\t\t[48]int8" + lines[idx].lstrip())
        idx += 1
        print(lines[idx])
        idx += 1
        print("\t\t\t[10]Column" + lines[idx].lstrip())
        idx += 1

        for i in range(10):
            print("\t\t\t\tColumn{[32]uint8" + lines[idx].lstrip()[1:])
            idx += 1
        
        print(lines[idx] + ",")
        idx += 1

        print(lines[idx])
        idx += 1
    
    print(lines[idx] + ",")
    idx += 1
    print(lines[idx][0])

    return idx + 1

def convert_lines(lines):
    idx = 0

    while (lines[idx] != "static song songdata = {"):
        idx += 1

    print("var SongData = Song{")
    print("\t[" + str(num_insts) + "]Instrument{ //Instruments")
    
    idx += 2

    incby = process_instruments(lines[idx:])
    idx += incby

    while 1:
        tokens = lines[idx].split(" ")
        idx += 1

        if tokens[0] == "#endif":
            break

        if tokens[0] != "#define":
            continue

        print("var " + tokens[1][4:] + " = " + tokens[2] + " " + ' '.join(tokens[3:]))

def main():
    print("package gonant\n")

    with open('gonantdata/music.h') as cmzk:
        lines = [line.rstrip() for line in cmzk.readlines()]

    convert_lines(lines)
        
if __name__ == "__main__":
    main()
