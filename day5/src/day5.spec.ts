
import { readFileInputByLines } from '../../reader/src/main';
import { calculateHighestTicketNumber, determineSeatLocation, getMissingTicketNumber } from './day5';
test('parse seat location properly', () => {
    const seatInput = 'FBFBBFFRLR';
    const seatLocation = determineSeatLocation(seatInput);
    expect(seatLocation.row).toBe(44);
    expect(seatLocation.column).toBe(5);
    expect(seatLocation.seatId).toBe(357);
})

describe('test some tickets', () => {
    test.each([
        ['BFFFBBFRRR', 70, 7, 567],
        ['FFFBBBFRRR', 14, 7, 119],
        ['BBFFBBFRLL', 102, 4, 820],
    ])
        ('test %s has %i rows %i cols and %i seat number',
            (seatInput: string, rowNum: number, colNum: number, seatNum: number) => {
                const seatLocation = determineSeatLocation(seatInput);
                expect(seatLocation.row).toBe(rowNum);
                expect(seatLocation.column).toBe(colNum);
                expect(seatLocation.seatId).toBe(seatNum);
            }
        )
})

test('get highest ticket number', () => {
    const tickets = ['BFFFBBFRRR', 'BBFFBBFRLL', 'FFFBBBFRRR'];
    const highestTicketNumber = calculateHighestTicketNumber(tickets);
    expect(highestTicketNumber).toBe(820)
})


test('calculate highest ticket numbers', () => {
    const tickets = readFileInputByLines('day5');
    const highestTicketNumber = calculateHighestTicketNumber(tickets);
    expect(highestTicketNumber).toBe(826)
})

test('get missing ticket number', () => {
    const tickets = readFileInputByLines('day5');
    const missingTicketNumber = getMissingTicketNumber(tickets);
    expect(missingTicketNumber).toBe(678)
})