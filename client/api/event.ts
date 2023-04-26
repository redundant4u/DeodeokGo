import { Event } from "@/types/event";
import http from "./index";

export const getEvents = async (): Promise<Event[]> => {
    return http.get(`events`);
};
