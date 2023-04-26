import { Event } from "@/types/event";
import Item from "./Item";

type PropTypes = {
    events: Event[];
    onEventBooked: (e: Event) => any;
};

const Body = ({ events, onEventBooked }: PropTypes) => {
    const items = events.map((event) => <Item key={event.ID} event={event} onBooked={() => onEventBooked(event)} />);

    return (
        <table className="table">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Where</th>
                    <th colSpan={2}>When (start/end)</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>{items}</tbody>
        </table>
    );
};

export default Body;
