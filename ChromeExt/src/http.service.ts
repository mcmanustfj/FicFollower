import { Story } from "./models";
import { Site } from "./models";
const SERVER_URL = 'http://localhost:5000';

function updateStory(story: Story) {
    return fetch(`${SERVER_URL}/story`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(story)
    });
}