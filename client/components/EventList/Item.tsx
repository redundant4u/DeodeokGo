import { Event } from "@/types/event";
import Link from "next/link";

type PropTypes = {
    event: Event;
    selected?: boolean;
    onBooked: () => any;
};

const Item = ({ event, selected, onBooked }: PropTypes) => {
    const start = new Date(event.StartDate * 1000);
    const end = new Date(event.EndDate * 1000);

    const locationName = event.Location ? event.Location.Name : "unknown";

    return (
        <tr>
            <td>{event.Name}</td>
            <td>{locationName}</td>
            <td>{start.toLocaleDateString()}</td>
            <td>{end.toLocaleDateString()}</td>
            <td>
                <Link href={""} />
                Book Now
            </td>
        </tr>
    );
};

export default Item;
