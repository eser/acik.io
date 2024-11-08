"use server";

import type { FormState, FormStateEntry } from "./types.ts";

const sendMessage = async (message: string) => {
  await fetch("http://localhost:8080/send", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      channelId: "1",
      message: {
        body: message
      }
    }),
  });
};

export const sendMessageAction = async (
  prevState: FormState,
  formData: FormData,
): Promise<FormState> => {
  const message = formData.get("message") as string;

  const start = performance.now();
  const _response = await sendMessage(message);
  const end = performance.now();
  const took = `${(end - start).toFixed(2)}ms`;

  const newMessage: FormStateEntry = [new Date(), message, took];

  return [...prevState, newMessage];
};
