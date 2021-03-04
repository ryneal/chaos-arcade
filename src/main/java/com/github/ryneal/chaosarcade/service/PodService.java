package com.github.ryneal.chaosarcade.service;

import com.github.ryneal.chaosarcade.model.Pod;

import java.util.List;
import java.util.Optional;

public interface PodService {

    Optional<Pod> getRandomPod();
    List<Pod> getAllPods();
    Optional<Pod> deletePod(Pod pod);

}
