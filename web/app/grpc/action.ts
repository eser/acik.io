"use server";

import * as grpc from "@grpc/grpc-js";
import * as broadcast from "../../proto/broadcast/broadcast.ts";
import type { FormState, FormStateEntry } from "./types.ts";

const sendMessage = (message: string) => {
  return new Promise((resolve, reject) => {
    const client = new broadcast.MessageServiceClient(
      "localhost:9090",
      grpc.credentials.createInsecure(),
    );

    client.send(
      {
        channelId: "1",
        message: {
          body: message,
        },
      },
      (error, response) => {
        if (error) {
          reject(error);
        } else {
          resolve(response);
        }
      },
    );
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
