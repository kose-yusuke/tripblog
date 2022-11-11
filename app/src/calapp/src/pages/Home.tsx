import { useCallback, useState, useEffect } from "react";
import FullCalendar, {
    DateSelectArg,
    EventApi,
    EventClickArg,
    EventInput
} from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import allLocales from "@fullcalendar/core/locales-all";
import interactionPlugin from "@fullcalendar/interaction";
import { INITIAL_EVENTS, createEventId } from "../event-utils";
import axios from 'axios'


const HomePage = () => {
    const [currentEvents, setCurrentEvents] = useState<EventInput[]>([]);
    // const handleEvents = useCallback(
    //     (events: EventApi[]) => setCurrentEvents(events),
    //     []
    // );
    const handleDateSelect = useCallback((selectInfo: DateSelectArg) => {
        let title = prompt("イベントのタイトルを入力してください")?.trim();
        let calendarApi = selectInfo.view.calendar;
        calendarApi.unselect();
        if (title) {
        calendarApi.addEvent({
            id: createEventId(),
            title,
            start: selectInfo.startStr,
            end: selectInfo.endStr,
            allDay: selectInfo.allDay
        });
        }
    }, []);
    const handleEventClick = useCallback((clickInfo: EventClickArg) => {
        if (
        window.confirm(`このイベント「${clickInfo.event.title}」を削除しますか`)
        ) {
        clickInfo.event.remove();
        
        }
    }, []);

    ///データ読み込み(一旦usersテーブルを読み込んでいる)
    useEffect(() => {
        async function fetchData() {
            const response = await axios.get('plans')
            setCurrentEvents(response.data)
        }

        fetchData();
    }, []);

    console.log(currentEvents)

    return(
    <div className="demo-app">
        <div className="demo-app-main">
        <FullCalendar
            plugins={[dayGridPlugin, interactionPlugin]}
            initialView="dayGridMonth"
            selectable={true}
            editable={true}
            initialEvents={INITIAL_EVENTS}
            events={currentEvents}
            locales={allLocales}
            locale="en"
            // eventsSet={handleEvents}
            select={handleDateSelect}
            eventClick={handleEventClick}
        />
        </div>
    </div>
    )
}

export default HomePage