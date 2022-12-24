package playground;

public class RailCompany {
	private Bookings bookings;
	private Schedule schedule;

	// Connection is previously reserved. Passenger needs to pay.

	/**
	 * A method to issue a previously reserved booking.
	 * 
	 * @param passenger  the passenger object which contains the passenger's
	 *                   infomation
	 * @param payment    payment object which contains payment information and
	 *                   methods that handles the payment
	 * @param connection a connection object which contains infomation of reserved
	 *                   trip schedule
	 * @return true if the booking is successfully issued, false if the booking
	 *         fails
	 */
	public boolean issueBooking(Passenger passenger, Payment payment, Connection connection) {
		// Get the newest record of the previously reserved Connection
		Connection reservedConnection = schedule.get(connection);
		if (!(reservedConnection.isReservationValid() && reservedConnection.isReservedBy(passenger))) {
			return false;
		}
		if (!payment.pay()) {
			return false;
		}
		reservedConnection.markAsPaid();
		schedule.save(reservedConnection);
		bookings.makeNewBooking(reservedConnection);
		return true;
	}

	// Connection is not previously reserved. No payment needed.

	/**
	 * A method to check and issue a booking for a passenger.
	 * 
	 * @param passenger   a Passenger object that contains the passenger's
	 *                    infomation
	 * @param date        date of the trip
	 * @param origin      origin of trip
	 * @param destination destination of the trip
	 * @return true if the booking is successfully issued, false if the booking
	 *         fails
	 */
	public boolean issueBooking(Passenger passenger, Date date, String origin, String destination) {
		// Check if there's an available connection.
		Connection connection = schedule.get(date, origin, destination);
		if (connection == null) {
			// If there's no connection for the specified date, origin and destination, the
			// booking fails, return false.
			return false;
		}
		// Available connection is found. Make a new booking for the passenger and
		// connection and return true.
		bookings.makeNewBooking(passenger, connection);
		return true;
	}

	// Connection is not previously reserved. No payment is needed. Revised version.

	/**
	 * A method to check and issue a booking for a passenger.
	 * 
	 * @param passenger a Passenger object that contains the passenger's
	 *                  information
	 * @param conn      the Connection reserved for the booking
	 * @return true if the booking is successfully issued, false if the booking
	 *         fails
	 */
	public boolean issueBooking(Passenger passenger, Connection conn) {
		if (conn == null) {
			// If there's no connection for the specified date, origin and destination, the
			// booking fails, return false.
			return false;
		}
		// Available connection is found. Make a new booking for the passenger and
		// connection. Return true if the booking is successfully made.
		return bookings.makeNewBooking(passenger, conn);
	}
}


1.boolean isSuccessful = terminal.makeBooking(passenger, date, origin ,destination){
  2.Connection conn = railCompany.get(passenger, date, origin, destination){
    3.Connection conn = schedule.getConnection(date, origin, destination){
      for(Connection conn in connList){
        4.boolean canSatisify = conn.checkAndMarkAsBooked(date, origin, destination);
        if(canSatisify) return conn;
      }
      return null;
    }
    5.boolean issueSuccessful = issueBooking(passenger, conn){
      if(conn == null) return false;
      6. boolean bookingSuccessful = makeNewBooking(passenger, conn){
        try {
          7. Ticket t = new Ticket(passenger, connection);
          tickets.add(t);
          return true;
        } catch (Exception e) {
          System.out.println("Error in creating ticket");
          return false;
        }
      }
      return bookingSuccessful;
    }
    if(bookingSuccessful) return conn;
    else return null;
  }
  return conn!=null;
}

List<Ticket> tickets = new ArrayList<Ticket>();

/**
 * A method to make new Booking
 * @param passenger A Passenger object that contains the passenger's information
 * @param connection The Connection to create ticket for
 * @return true if the a Ticket is successfully created, false if it fails
 */
public boolean makeNewBooking(Passenger passenger, Connection connection) {
		try {
				//Creates an object of ticket
				Ticket ticket = new Ticket(passenger, connection);
				//add ticket to list of tickets
				tickets.add(ticket);
				//if it is created successfully then return true else false
				return true;
		} catch (Exception e) {
				System.out.println("Error in creating ticket");
				return false;
		}
}
