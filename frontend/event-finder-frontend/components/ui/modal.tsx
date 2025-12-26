"use client"

import React, { useEffect } from "react";

type ModalProps = {
    title: string;
    isOpen: boolean;
    onClose: () => void;
    children: React.ReactNode;
};

const Modal = ({ title, isOpen, onClose, children }: ModalProps) => {
    useEffect(() => {
        const handleEscape = (e: KeyboardEvent) => {
            if (e.key === "Escape" && isOpen) {
                onClose();
            }
        };

        document.addEventListener("keydown", handleEscape);
        return () => document.removeEventListener("keydown", handleEscape);
    }, [isOpen, onClose]);

    if (!isOpen) return null;

    return (
        <div
            className="fixed inset-0 z-50 bg-black/50 backdrop-blur-sm flex items-center justify-center px-4"
            onClick={onClose}
        >
            <div
                className="relative bg-background w-full max-w-md rounded-xl shadow-lg p-6 animate-fade-in-up"
                onClick={(e) => e.stopPropagation()}
            >
                <button
                    onClick={onClose}
                    className="absolute top-4 right-4 text-muted-foreground hover:text-destructive text-xl font-bold transition-colors"
                    aria-label="Close"
                >
                    ×
                </button>

                <h2 className="text-2xl font-bold text-foreground mb-4 text-center">
                    {title}
                </h2>

                <div>{children}</div>
            </div>
        </div>
    );
};

export default Modal;
