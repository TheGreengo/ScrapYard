def getMpg(odd1, odd2, gallons):
    return (odd1 - odd2) / gallons

def getKpL(mpg):
    return (mpg * 1.6) / 3.8

def main():
    start = float(input("oddometer start"))
    stop = float(input("oddometer stop"))
    used = float(input("gallons used"))
    mpg = getMpg(start, stop, used)
    kpl = getKpL(mpg)
    print(f"You got {mpg} miles per gallon and {kpl} kilometers per liter")
    return 0

main()