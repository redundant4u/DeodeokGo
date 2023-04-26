import { Event } from "@/types/event";
import Loader from "@/components/Loader";
import Body from "./Body";

type PropTypes = {
    events: Event[];
};

const EventList = ({ events }: PropTypes) => {
    const handleEventBooked = (e: Event) => {
        console.log("booking event");
    };

    return <Body events={events} onEventBooked={(e) => handleEventBooked(e)} />;
};

export default EventList;
