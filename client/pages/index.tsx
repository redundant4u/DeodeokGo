import { GetStaticProps, NextPage } from "next";

import EventList from "@/components/EventList";
import Navigation from "@/components/Navigation";
import { getEvents } from "@/api/event";
import { Event } from "@/types/event";

type PropTypes = {
    events: Event[];
};

const Home: NextPage<PropTypes> = ({ events }) => {
    return (
        <>
            <Navigation brandName="MyEvents" />
            <div>
                <h1>My Events</h1>
                <EventList events={events} />
            </div>
        </>
    );
};

export const getStaticProps: GetStaticProps = async () => {
    const events = await getEvents();

    return {
        props: { events },
    };
};

export default Home;
