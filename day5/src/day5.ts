function convertToBinaryAndGetNumber(inputStr: string) {
    // Represent bits, F is 0, B is 1
    // L i 0, R is 1
    const binaryRep = inputStr.split('')
        .map(c => ["R", "B"].includes(c) ? "1" : "0").join("");
    return parseInt(binaryRep, 2)
}

export function determineSeatLocation(seatInput: string): SeatLocation {
    const rowLocation = seatInput.substring(0, 7)
    const columnLocation = seatInput.substring(7, 10)
    const row = convertToBinaryAndGetNumber(rowLocation)
    const column = convertToBinaryAndGetNumber(columnLocation)
    return {
        seatInput,
        row,
        column,
        seatId: row * 8 + column,
    }
}

export function calculateHighestTicketNumber(seatInputs: string[]): number {
    const ticketInfo = seatInputs.map((seatInput: string) => determineSeatLocation(seatInput))
    const seatNumbers = ticketInfo.map(s => s.seatId)
    return Math.max(...seatNumbers)
}

export function getMissingTicketNumber(seatInputs: string[]): number {
    const ticketInfo = seatInputs.map((seatInput: string) => determineSeatLocation(seatInput))
    const sortedSeatNumbers = ticketInfo.map(s => s.seatId).sort((a, b) => a - b);
    const seatNumberSet = new Set(sortedSeatNumbers);
    for(let i = sortedSeatNumbers[0]; i < sortedSeatNumbers[sortedSeatNumbers.length - 1]; i++){
        if(!seatNumberSet.has(i) && (seatNumberSet.has(i-1) || seatNumberSet.has(i-1))){
            return i;
        }
    }
    throw new Error("No seat number found")
}

interface SeatLocation {
    seatInput: string,
    row: number,
    column: number,
    seatId: number,
}