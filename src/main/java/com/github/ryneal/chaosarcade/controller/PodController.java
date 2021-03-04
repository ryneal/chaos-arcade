package com.github.ryneal.chaosarcade.controller;

import com.github.ryneal.chaosarcade.model.Pod;
import com.github.ryneal.chaosarcade.service.PodService;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class PodController {

    private final PodService podService;

    public PodController(PodService podService) {
        this.podService = podService;
    }

    @GetMapping("/pods")
    public ResponseEntity<List<Pod>> getAllPods() {
        return ResponseEntity.ok(this.podService.getAllPods());
    }

    @DeleteMapping("/pods/random")
    public ResponseEntity<Pod> deleteRandomPod() {
        return ResponseEntity.of(this.podService.getRandomPod().flatMap(this.podService::deletePod));
    }

}
